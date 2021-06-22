package backend

import (
	"context"
	"fmt"

	"github.com/VGLoic/godel/ethshell"
	"github.com/VGLoic/godel/eventlog"
	"github.com/VGLoic/godel/ipfsshell"
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

func (b *Backend) PublishEvent(ctx context.Context, publishRequest PublishRequest) (eventlog.Event, error) {
	businessEvent := ipfsshell.BusinessEvent{
		Type:    publishRequest.Type,
		Payload: publishRequest.Payload,
		Version: publishRequest.Version,
	}
	cid, ipfsPublishErr := b.ipfs.PublishBusinessEvent(businessEvent)
	if ipfsPublishErr != nil {
		return eventlog.Event{}, ipfsPublishErr
	}

	tx, from, txErr := b.eth.PublishEvent(ctx, publishRequest.Topic, cid, publishRequest.NewAccounts)
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
		newAccounts := []string{}
		for _, newAccount := range event.NewAccounts {
			newAccounts = append(newAccounts, newAccount.String())
		}
		completedEvent := eventlog.Event{
			Type:        businessEvent.Type,
			Payload:     businessEvent.Payload,
			Version:     businessEvent.Version,
			Topic:       topic,
			Cid:         event.Cid,
			Emitter:     event.Emitter.String(),
			NewAccounts: newAccounts,
			Timestamp:   event.Timestamp.Uint64(),
			BlockNumber: event.BlockNumber.Uint64(),
		}
		completedEvents = append(completedEvents, completedEvent)
	}
	return completedEvents, nil
}

func (b *Backend) FindPendingEvents() ([]eventlog.Event, error) {
	return b.eventLog.FindPendingEvents()
}

func (b *Backend) FindConfirmedEvents() ([]eventlog.Event, error) {
	return b.eventLog.FindConfirmedEvents()
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
				topic, cid, newAccounts, unpackingErr := ethshell.UnpackLog(log)
				if unpackingErr != nil {
					fmt.Println(fmt.Errorf("Error in unpacking of the log: %s \n", unpackingErr))
					break
				}
				emitter, getEmitterErr := b.eth.GetTxSender(ctx, log.TxHash)
				if getEmitterErr != nil {
					fmt.Println(fmt.Errorf("Error in getting the tx sender: %s \n", getEmitterErr))
					break
				}
				blockHeader, blockHeaderErr := b.eth.HeaderByNumber(ctx, log.BlockNumber)
				if blockHeaderErr != nil {
					fmt.Println(fmt.Errorf("Error in getting the block header: %s \n", blockHeaderErr))
					break
				}

				if emitter == b.accountAddress {
					_, confirmationErr := b.eventLog.Confirm(log.TxHash, log.BlockNumber, blockHeader.Time)
					if confirmationErr != nil {
						fmt.Println(fmt.Errorf("Error in confirming event: %s \n", confirmationErr))
						break
					}
				} else {
					// If topic is known, add event
					lastBlockNumber, lastBlockErr := b.eventLog.FindLastSynchronisedBlockNumber(topic)
					if lastBlockErr != nil {
						fmt.Println(fmt.Errorf("Error in finding last synchronised block: %s \n", lastBlockErr))
						break
					}
					if isTopicKnown := lastBlockNumber > 0; isTopicKnown {
						businessEvent, retrieveFromIpfsErr := b.ipfs.GetBusinessEvent(cid)
						if retrieveFromIpfsErr != nil {
							fmt.Println(fmt.Errorf("Error in retrieving the informations from IPFS: %s \n", retrieveFromIpfsErr))
							break
						}
						event := eventlog.Event{
							Type:        businessEvent.Type,
							Payload:     businessEvent.Payload,
							Version:     businessEvent.Version,
							Topic:       topic,
							Cid:         cid,
							Emitter:     emitter,
							NewAccounts: newAccounts,
							Timestamp:   blockHeader.Time,
							BlockNumber: log.BlockNumber,
							TxHash:      log.TxHash.String(),
						}
						_, insertionErr := b.eventLog.Insert(event)
						if insertionErr != nil {
							fmt.Println(fmt.Errorf("Error in inserting new event: %s \n", insertionErr))
							break
						}
					} else {
						if isMemberAdded := contains(newAccounts, b.accountAddress); isMemberAdded {
							syncErr := b.synchroniseTopic(ctx, topic)
							if syncErr != nil {
								fmt.Println(fmt.Errorf("Error in synchronising topic: %s \n", syncErr))
								break
							}
						} else {
							fmt.Println("Just looking at an event that is not for me :(")
						}
					}
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

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
