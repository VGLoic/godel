package backend

import (
	"context"
	"fmt"

	"github.com/VGLoic/godel/eventlog"
	"github.com/VGLoic/godel/ipfsshell"
	"github.com/ethereum/go-ethereum/core/types"
)

type Backend struct {
	eventLog       EventLog
	ipfs           IpfsShell
	eth            EthShell
	accountAddress string
}

func NewBackend(eventLog EventLog, ipfs IpfsShell, eth EthShell, accountAddress string) *Backend {
	return &Backend{
		eventLog:       eventLog,
		ipfs:           ipfs,
		eth:            eth,
		accountAddress: accountAddress,
	}
}

type PublishRequest struct {
	Type        string
	Payload     string
	Version     float64
	Topic       string
	NewAccounts []string
}

func (b *Backend) PublishEvent(publishRequest PublishRequest) (eventlog.Event, error) {
	businessEvent := ipfsshell.BusinessEvent{
		Type:    publishRequest.Type,
		Payload: publishRequest.Payload,
		Version: publishRequest.Version,
	}
	cid, ipfsPublishErr := b.ipfs.PublishBusinessEvent(businessEvent)
	if ipfsPublishErr != nil {
		return eventlog.Event{}, ipfsPublishErr
	}

	tx, from, txErr := b.eth.PublishEvent(context.Background(), publishRequest.Topic, cid, publishRequest.NewAccounts)
	if txErr != nil {
		return eventlog.Event{}, txErr
	}

	event := eventlog.Event{
		Type:        businessEvent.Type,
		Payload:     businessEvent.Payload,
		Version:     businessEvent.Version,
		Topic:       publishRequest.Topic,
		Cid:         cid,
		Emitter:     from.String(),
		NewAccounts: publishRequest.NewAccounts,
		Timestamp:   0,
		BlockNumber: 0,
		TxHash:      tx.Hash().String(),
	}
	insertedEvent, insertionErr := b.eventLog.Insert(event)
	if insertionErr != nil {
		return eventlog.Event{}, insertionErr
	}
	return insertedEvent, nil
}

func (b *Backend) FindPendingEvents() ([]eventlog.Event, error) {
	return b.eventLog.FindPendingEvents()
}

func (b *Backend) FindConfirmedEvents() ([]eventlog.Event, error) {
	return b.eventLog.FindConfirmedEvents()
}

func (b *Backend) SynchroniseAllTopics(ctx context.Context) error {
	topics, topicsErr := b.eth.GetTopics(ctx)
	if topicsErr != nil {
		return topicsErr
	}

	for _, topic := range topics {
		synchronisationErr := b.synchroniseTopic(ctx, topic)
		if synchronisationErr != nil {
			return synchronisationErr
		}
	}

	return nil
}

func (b *Backend) synchroniseTopic(ctx context.Context, topic string) error {

	lastBlocknumber, blockNumberErr := b.eventLog.FindLastSynchronisedBlockNumber(topic)
	if blockNumberErr != nil {
		return blockNumberErr
	}
	fmt.Println("Last block: ", lastBlocknumber, " for topic ", topic)
	events, eventsErr := b.getEvents(topic, lastBlocknumber)
	if eventsErr != nil {
		return eventsErr
	}

	clearErr := b.eventLog.ClearPendingEvents(topic)
	if clearErr != nil {
		return clearErr
	}

	resetErr := b.eventLog.ResetFromBlockAndInsert(topic, lastBlocknumber, events)
	if resetErr != nil {
		return resetErr
	}

	return nil
}

func (b *Backend) getEvents(topic string, fromBlock uint64) ([]eventlog.Event, error) {
	events, eventsErr := b.eth.GetEvents(topic, fromBlock)
	if eventsErr != nil {
		return nil, eventsErr
	}
	completedEvents := []eventlog.Event{}
	for _, event := range events {
		businessEvent, businessEventErr := b.ipfs.GetBusinessEvent(event.Cid)
		if businessEventErr != nil {
			return nil, businessEventErr
		}
		pinErr := b.ipfs.PinBusinessEventCid(event.Cid)
		if pinErr != nil {
			return nil, pinErr
		}
		completedEvent := eventlog.Event{
			Type:        businessEvent.Type,
			Payload:     businessEvent.Payload,
			Version:     businessEvent.Version,
			Topic:       topic,
			Cid:         event.Cid,
			Emitter:     event.Emitter,
			NewAccounts: event.NewAccounts,
			Timestamp:   event.Timestamp,
			BlockNumber: event.BlockNumber,
		}
		completedEvents = append(completedEvents, completedEvent)
	}
	return completedEvents, nil
}

func (b *Backend) SubscribeToEvents(ctx context.Context) error {
	sub, logs, subErr := b.eth.SubscribeToEvents(ctx)
	if subErr != nil {
		return subErr
	}

	go func() {
		for {
			select {
			case log := <-logs:
				processErr := b.processLog(ctx, log)
				if processErr != nil {
					fmt.Println(fmt.Errorf("Error in processing the log: %s \n", processErr))
					break
				}
			case err := <-sub.Err():
				fmt.Println(fmt.Errorf("Error in Ethereum Subscription: %s \n", err))
				close(logs)
				return
			case <-ctx.Done():
				close(logs)
				sub.Unsubscribe()
				return
			}
		}
	}()

	return nil
}

func (b *Backend) processLog(ctx context.Context, log types.Log) error {

	rawEvent, unpackingErr := b.eth.UnpackLog(ctx, log)
	if unpackingErr != nil {
		return unpackingErr
	}

	if rawEvent.Emitter == b.accountAddress {
		_, confirmationErr := b.eventLog.Confirm(rawEvent.TxHash, rawEvent.BlockNumber, rawEvent.Timestamp)
		if confirmationErr != nil {
			return confirmationErr
		}
		return nil
	}

	lastBlockNumber, lastBlockErr := b.eventLog.FindLastSynchronisedBlockNumber(rawEvent.TopicId)
	if lastBlockErr != nil {
		return lastBlockErr
	}
	if isTopicKnown := lastBlockNumber > 0; isTopicKnown {
		businessEvent, retrieveFromIpfsErr := b.ipfs.GetBusinessEvent(rawEvent.Cid)
		if retrieveFromIpfsErr != nil {
			return retrieveFromIpfsErr
		}
		pinErr := b.ipfs.PinBusinessEventCid(rawEvent.Cid)
		if pinErr != nil {
			return pinErr
		}
		event := eventlog.Event{
			Type:        businessEvent.Type,
			Payload:     businessEvent.Payload,
			Version:     businessEvent.Version,
			Topic:       rawEvent.TopicId,
			Cid:         rawEvent.Cid,
			Emitter:     rawEvent.Emitter,
			NewAccounts: rawEvent.NewAccounts,
			Timestamp:   rawEvent.Timestamp,
			BlockNumber: rawEvent.BlockNumber,
			TxHash:      rawEvent.TxHash,
		}
		_, insertionErr := b.eventLog.Insert(event)
		if insertionErr != nil {
			return insertionErr
		}
		return nil
	}
	if isMemberAdded := contains(rawEvent.NewAccounts, b.accountAddress); isMemberAdded {
		syncErr := b.synchroniseTopic(ctx, rawEvent.TopicId)
		if syncErr != nil {
			return syncErr
		}
		return nil
	}

	fmt.Println("Just looking at an event that is not for me :(")
	return nil
}

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// #################################### WIP ####################################
// func (b *Backend) LaunchPinningRoutine(ctx context.Context) error {
// 	// go func() {
// 	// 	next := time.After(0)
// 	// 	for {
// 	// 		select {
// 	// 		case <-next:
// 	// 			fmt.Println("Triggered")
// 	// 			// b.parseEvents(ctx)
// 	// 			next = time.After(10 * time.Second)
// 	// 		case <-ctx.Done():
// 	// 			return
// 	// 		}
// 	// 	}
// 	// }()
// 	return nil
// }

// func (b *Backend) parseEvents(ctx context.Context) {
// 	pageSize := uint(1)
// 	page := uint(2)

// 	hasReachedLastRow := false
// 	for !hasReachedLastRow {
// 		fmt.Println("Dealing with page: ", page)
// 		offset := (page - 1) * pageSize
// 		hasReachedLastRow = b.parsePage(ctx, offset, pageSize)
// 		page += 1
// 	}
// }

// func (b *Backend) parsePage(ctx context.Context, offset uint, pageSize uint) bool {
// 	events, queryErr := b.eventLog.FindPage(offset, pageSize)
// 	if queryErr != nil {
// 		fmt.Println(fmt.Errorf("Error in querying events for offset %v and pageSize %v: %s \n", offset, pageSize, queryErr))
// 		return false
// 	}
// 	for _, event := range events {
// 		fmt.Println("Event: ", event)
// 	}

// 	return len(events) < int(pageSize)
// }
