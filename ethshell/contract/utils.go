package contract

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func UnpackLog(log types.Log) (ContractEventPublished, error) {
	contractAbi, abiErr := abi.JSON(strings.NewReader(string(ContractABI)))
	if abiErr != nil {
		return ContractEventPublished{}, abiErr
	}

	unpacked, unpackErr := contractAbi.Unpack("EventPublished", log.Data)

	if unpackErr != nil {
		return ContractEventPublished{}, unpackErr
	}
	if len(unpacked) < 3 {
		return ContractEventPublished{}, errors.New("Oh no, the unpacked event must have four elements!")
	}
	topic := unpacked[0].(string)
	cid := unpacked[1].(string)
	newAccounts := unpacked[2].([]common.Address)

	event := ContractEventPublished{
		TopicId:     topic,
		Cid:         cid,
		NewAccounts: newAccounts,
		Raw:         log,
	}

	return event, nil
}
