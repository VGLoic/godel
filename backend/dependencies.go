package backend

import (
	"context"

	"github.com/VGLoic/godel/ethshell"
	"github.com/VGLoic/godel/eventlog"
	"github.com/VGLoic/godel/ipfsshell"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EventLog interface {
	Insert(event eventlog.Event) (eventlog.Event, error)
	ResetFromBlockAndInsert(topic string, lastBlockNumber uint64, events []eventlog.Event) error
	FindLastSynchronisedBlockNumber(topic string) (uint64, error)
	FindPendingEvents() ([]eventlog.Event, error)
	FindConfirmedEvents() ([]eventlog.Event, error)
	Confirm(txHash common.Hash, blockNumber uint64, timestamp uint64) (eventlog.Event, error)
	ClearPendingEvents(topic string) error
}
type IpfsShell interface {
	PublishBusinessEvent(businessEvent ipfsshell.BusinessEvent) (string, error)
	GetBusinessEvent(cid string) (ipfsshell.BusinessEvent, error)
}
type EthShell interface {
	PublishEvent(ctx context.Context, topic string, cid string, newAccounts []string) (*types.Transaction, common.Address, error)
	GetTopics(ctx context.Context) ([]string, error)
	GetEvents(topic string, fromBlock uint64) ([]ethshell.EventLogEvent, error)
	SubscribeToEvents(ctx context.Context) (ethereum.Subscription, chan types.Log, error)
	GetTxSender(ctx context.Context, txHash common.Hash) (string, error)
	HeaderByNumber(ctx context.Context, blockNumber uint64) (*types.Header, error)
}
