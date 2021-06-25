package ethshell

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Shell struct {
	ethClient        *ethclient.Client
	contractInstance *Del
	privateKey       string
	contractAddress  string
}

type ShellConfiguration struct {
	EthNodeUrl      string
	ContractAddress string
	PrivateKey      string
}

func NewShell(shellConfiguration ShellConfiguration) (*Shell, error) {
	ethClient, ethClientErr := ethclient.Dial(shellConfiguration.EthNodeUrl)
	if ethClientErr != nil {
		return nil, ethClientErr
	}

	address := common.HexToAddress(shellConfiguration.ContractAddress)
	contractInstance, contractErr := NewDel(address, ethClient)
	if contractErr != nil {
		return nil, contractErr
	}

	ethShell := Shell{
		ethClient:        ethClient,
		contractInstance: contractInstance,
		privateKey:       shellConfiguration.PrivateKey,
		contractAddress:  shellConfiguration.ContractAddress,
	}

	return &ethShell, nil
}

func (s *Shell) SubscribeToEvents(ctx context.Context) (ethereum.Subscription, chan types.Log, error) {
	logs := make(chan types.Log)
	contractAddress := common.HexToAddress(s.contractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	sub, err := s.ethClient.SubscribeFilterLogs(ctx, query, logs)

	if err != nil {
		return nil, nil, err
	}

	return sub, logs, nil
}

func (s *Shell) GetTopics(ctx context.Context) ([]string, error) {
	privateKey, pkErr := crypto.HexToECDSA(s.privateKey)
	if pkErr != nil {
		return nil, pkErr
	}

	fromAddress, fromAddressErr := deriveFromAddress(privateKey)
	if fromAddressErr != nil {
		return nil, fromAddressErr
	}

	topics, getTopicsErr := s.contractInstance.GetTopics(nil, fromAddress)
	if getTopicsErr != nil {
		return nil, getTopicsErr
	}

	return topics, nil
}

func (s *Shell) GetEvents(topic string, fromBlock uint64) ([]EventLogEvent, error) {
	events, eventsErr := s.contractInstance.GetEvents(nil, topic)
	if eventsErr != nil {
		return nil, eventsErr
	}

	filteredEvents := []EventLogEvent{}
	for _, event := range events {
		if event.BlockNumber.Uint64() >= fromBlock {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents, nil
}

func (s *Shell) PublishEvent(ctx context.Context, topic string, cid string, newAccounts []string) (*types.Transaction, common.Address, error) {

	txOptions, from, txOptionsErr := s.deriveTransactionOpts(ctx)
	if txOptionsErr != nil {
		return nil, common.Address{}, txOptionsErr
	}

	newAccountAddresses := []common.Address{}
	for _, newAccount := range newAccounts {
		newAccountAddresses = append(newAccountAddresses, common.HexToAddress(newAccount))
	}
	tx, txErr := s.contractInstance.PublishEvent(txOptions, topic, cid, newAccountAddresses)
	if txErr != nil {
		return nil, common.Address{}, txErr
	}

	return tx, from, nil
}

func UnpackLog(log types.Log) (string, string, []string, error) {
	contractAbi, abiErr := abi.JSON(strings.NewReader(string(DelABI)))
	if abiErr != nil {
		return "", "", []string{}, abiErr
	}

	unpacked, unpackErr := contractAbi.Unpack("EventPublished", log.Data)

	if unpackErr != nil {
		return "", "", []string{}, unpackErr
	}
	if len(unpacked) < 3 {
		return "", "", []string{}, errors.New("Oh no, the unpacked event must have three elements!")
	}
	topic := unpacked[0].(string)
	cid := unpacked[1].(string)
	newAddresses := unpacked[2].([]common.Address)
	newAccounts := []string{}
	for _, newAddress := range newAddresses {
		newAccounts = append(newAccounts, newAddress.String())
	}

	return topic, cid, newAccounts, nil
}

func deriveFromAddress(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return common.Address{}, errors.New("Oh no, no public key ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress, nil
}

func (s *Shell) deriveTransactionOpts(ctx context.Context) (*bind.TransactOpts, common.Address, error) {
	privateKey, pkErr := crypto.HexToECDSA(s.privateKey)
	if pkErr != nil {
		return nil, common.Address{}, pkErr
	}

	fromAddress, fromAddressErr := deriveFromAddress(privateKey)
	if fromAddressErr != nil {
		return nil, common.Address{}, fromAddressErr
	}

	nonce, nonceErr := s.ethClient.PendingNonceAt(ctx, fromAddress)
	if nonceErr != nil {
		return nil, common.Address{}, nonceErr
	}
	gasPrice, gasErr := s.ethClient.SuggestGasPrice(ctx)
	if gasErr != nil {
		return nil, common.Address{}, gasErr
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	// Could be sent as nil, will be fixed in next release of geth --> see https://github.com/ethereum/go-ethereum/pull/23062
	auth.Context = ctx

	return auth, fromAddress, nil
}

func (s *Shell) GetTxSender(ctx context.Context, txHash common.Hash) (string, error) {
	tx, _, txErr := s.ethClient.TransactionByHash(ctx, txHash)
	if txErr != nil {
		return "", txErr
	}

	chainID, chainIDErr := s.ethClient.NetworkID(ctx)
	if chainIDErr != nil {
		return "", chainIDErr
	}

	msg, msgErr := tx.AsMessage(types.NewEIP155Signer(chainID), nil)
	if msgErr != nil {
		return "", msgErr
	}

	return msg.From().String(), nil
}

func (s *Shell) HeaderByNumber(ctx context.Context, blockNumber uint64) (*types.Header, error) {
	header, err := s.ethClient.HeaderByNumber(ctx, big.NewInt(int64(blockNumber)))
	return header, err
}
