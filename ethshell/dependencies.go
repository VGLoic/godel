package ethshell

import (
	"github.com/VGLoic/godel/ethshell/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EventLogContract interface {
	GetTopics(opts *bind.CallOpts, address common.Address) ([]string, error)
	GetEvents(opts *bind.CallOpts, topic string) ([]contract.EventLogEvent, error)
	PublishEvent(opts *bind.TransactOpts, topic string, cid string, newAccounts []common.Address) (*types.Transaction, error)
}
