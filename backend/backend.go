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
	cid, err := b.ipfs.PublishBusinessEvent(businessEvent)
	if err != nil {
		return eventlog.Event{}, err
	}

	tx, from, err := b.eth.PublishEvent(context.Background(), publishRequest.Topic, cid, publishRequest.NewAccounts)
	if err != nil {
		return eventlog.Event{}, err
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
		Depth:       0,
		TxHash:      tx.Hash().String(),
	}
	insertedEvent, err := b.eventLog.Insert(event)
	if err != nil {
		return eventlog.Event{}, err
	}
	return insertedEvent, nil
}

func (b *Backend) FindPendingEvents() ([]eventlog.Event, error) {
	return b.eventLog.FindPendingEvents()
}

func (b *Backend) FindConfirmedEvents() ([]eventlog.Event, error) {
	return b.eventLog.FindConfirmedEvents()
}

func (b *Backend) MakeLocalDataAvailable(ctx context.Context) error {
	pageSize := uint(50)
	page := uint(1)

	hasReachedLastRow := false
	var err error
	for !hasReachedLastRow {
		offset := (page - 1) * pageSize
		hasReachedLastRow, err = b.makePageAvailable(ctx, offset, pageSize)
		if err != nil {
			fmt.Println(fmt.Errorf("Error in making the page available for offset %v and pageSize %v: %s \n", offset, pageSize, err))
		}
	}
	return nil
}

func (b *Backend) makePageAvailable(ctx context.Context, offset uint, pageSize uint) (bool, error) {
	events, hasMore, err := b.eventLog.FindPage(offset, pageSize)
	if err != nil {
		fmt.Println(fmt.Errorf("Error in querying events for offset %v and pageSize %v: %s \n", offset, pageSize, err))
		return false, err
	}
	for _, event := range events {
		businessEvent := ipfsshell.BusinessEvent{
			Type:    event.Type,
			Payload: event.Payload,
			Version: event.Version,
		}
		_, publishErr := b.ipfs.PublishBusinessEvent(businessEvent)
		if publishErr != nil {
			fmt.Println(fmt.Errorf("Error in publishing event to IPFS during initialization: %s \n", publishErr))
		}
	}
	return !hasMore, nil
}

func (b *Backend) SynchroniseAllTopics(ctx context.Context) error {
	topics, err := b.eth.GetTopics(ctx)
	if err != nil {
		return err
	}

	for _, topic := range topics {
		err := b.synchroniseTopic(ctx, topic)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Backend) synchroniseTopic(ctx context.Context, topic string) error {

	lastDepth, err := b.eventLog.FindLastSynchronisedDepth(topic)
	if err != nil {
		return err
	}
	fmt.Println("Last depth: ", lastDepth, " for topic ", topic)
	events, err := b.getEvents(topic, lastDepth)
	if err != nil {
		return err
	}

	err = b.eventLog.ClearPendingEvents(topic)
	if err != nil {
		return err
	}

	_, err = b.eventLog.InsertMany(events)
	if err != nil {
		return err
	}

	return nil
}

func (b *Backend) getEvents(topic string, fromDepth uint64) ([]eventlog.Event, error) {
	events, err := b.eth.GetEvents(topic, fromDepth)
	if err != nil {
		return nil, err
	}
	completedEvents := []eventlog.Event{}
	for _, event := range events {
		businessEvent, businessEventErr := b.ipfs.GetBusinessEvent(event.Cid)
		if businessEventErr != nil {
			return nil, businessEventErr
		}
		err := b.ipfs.PinBusinessEventCid(event.Cid)
		if err != nil {
			return nil, err
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
			Depth:       event.Depth,
		}
		completedEvents = append(completedEvents, completedEvent)
	}
	return completedEvents, nil
}

func (b *Backend) SubscribeToEvents(ctx context.Context) error {
	sub, logs, err := b.eth.SubscribeToEvents(ctx)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case log := <-logs:
				err := b.processLog(ctx, log)
				if err != nil {
					fmt.Println(fmt.Errorf("Error in processing the log: %s \n", err))
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

	rawEvent, err := b.eth.UnpackLog(ctx, log)
	if err != nil {
		return err
	}

	if rawEvent.Emitter == b.accountAddress {
		_, err := b.eventLog.Confirm(rawEvent.TxHash, rawEvent.BlockNumber, rawEvent.Timestamp, rawEvent.Depth)
		if err != nil {
			return err
		}
		return nil
	}

	lastDepth, err := b.eventLog.FindLastSynchronisedDepth(rawEvent.TopicId)
	if err != nil {
		return err
	}
	if isTopicKnown := lastDepth > 0; isTopicKnown {
		businessEvent, err := b.ipfs.GetBusinessEvent(rawEvent.Cid)
		if err != nil {
			return err
		}
		err = b.ipfs.PinBusinessEventCid(rawEvent.Cid)
		if err != nil {
			return err
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
			Depth:       rawEvent.Depth,
			TxHash:      rawEvent.TxHash,
		}
		_, err = b.eventLog.Insert(event)
		if err != nil {
			return err
		}
		return nil
	}
	if isMemberAdded := contains(rawEvent.NewAccounts, b.accountAddress); isMemberAdded {
		err := b.synchroniseTopic(ctx, rawEvent.TopicId)
		if err != nil {
			return err
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
