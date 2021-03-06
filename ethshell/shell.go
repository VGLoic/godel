package ethshell

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/VGLoic/godel/ethshell/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EventLogEvent struct {
	TopicId     string
	Cid         string
	NewAccounts []string
	Emitter     string
	Timestamp   uint64
	BlockNumber uint64
	Depth       uint64
	TxHash      string
}

type Shell struct {
	ethClient        *ethclient.Client
	contractInstance EventLogContract
	privateKey       string
	contractAddress  string
}

type ShellConfiguration struct {
	EthNodeUrl      string
	ContractAddress string
	PrivateKey      string
}

func NewShell(shellConfiguration ShellConfiguration) (*Shell, error) {
	ethClient, err := ethclient.Dial(shellConfiguration.EthNodeUrl)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(shellConfiguration.ContractAddress)
	contractInstance, err := contract.NewContract(address, ethClient)
	if err != nil {
		return nil, err
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
	privateKey, err := crypto.HexToECDSA(s.privateKey)
	if err != nil {
		return nil, err
	}

	fromAddress, err := deriveFromAddress(privateKey)
	if err != nil {
		return nil, err
	}

	topics, err := s.contractInstance.GetTopics(nil, fromAddress)
	if err != nil {
		return nil, err
	}

	return topics, nil
}

func (s *Shell) GetEvents(topic string, fromDepth uint64) ([]EventLogEvent, error) {
	events, err := s.contractInstance.GetEvents(nil, topic)
	if err != nil {
		return nil, err
	}
	formattedEvents := []EventLogEvent{}
	for index, e := range events {
		depth := uint64(index) + 1
		if depth > fromDepth {
			newAccounts := []string{}
			for _, address := range e.NewAccounts {
				newAccounts = append(newAccounts, address.String())
			}
			formattedEvent := EventLogEvent{
				TopicId:     topic,
				Cid:         e.Cid,
				Emitter:     e.Emitter.String(),
				NewAccounts: newAccounts,
				Timestamp:   e.Timestamp.Uint64(),
				BlockNumber: e.BlockNumber.Uint64(),
				Depth:       depth,
			}
			formattedEvents = append(formattedEvents, formattedEvent)
		}
	}
	return formattedEvents, nil
}

func (s *Shell) PublishEvent(
	ctx context.Context,
	topic string,
	cid string,
	newAccounts []string,
) (*types.Transaction, common.Address, error) {

	txOptions, from, err := s.deriveTransactionOpts(ctx)
	if err != nil {
		return nil, common.Address{}, err
	}

	newAccountAddresses := []common.Address{}
	for _, newAccount := range newAccounts {
		newAccountAddresses = append(newAccountAddresses, common.HexToAddress(newAccount))
	}
	tx, err := s.contractInstance.PublishEvent(txOptions, topic, cid, newAccountAddresses)
	if err != nil {
		return nil, common.Address{}, err
	}

	return tx, from, nil
}

func (s *Shell) UnpackLog(ctx context.Context, log types.Log) (EventLogEvent, error) {
	rawEvent, err := contract.UnpackLog(log)
	if err != nil {
		return EventLogEvent{}, err
	}

	emitter, err := s.GetTxSender(ctx, log.TxHash)
	if err != nil {
		return EventLogEvent{}, err
	}
	header, err := s.ethClient.HeaderByNumber(ctx, big.NewInt(int64(log.BlockNumber)))
	if err != nil {
		return EventLogEvent{}, err
	}

	newAccounts := []string{}
	for _, address := range rawEvent.NewAccounts {
		newAccounts = append(newAccounts, address.String())
	}

	event := EventLogEvent{
		TopicId:     rawEvent.TopicId,
		Cid:         rawEvent.Cid,
		NewAccounts: newAccounts,
		Emitter:     emitter.String(),
		Timestamp:   header.Time,
		BlockNumber: header.Number.Uint64(),
		Depth:       rawEvent.Depth.Uint64(),
		TxHash:      rawEvent.Raw.TxHash.String(),
	}

	return event, nil
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
	privateKey, err := crypto.HexToECDSA(s.privateKey)
	if err != nil {
		return nil, common.Address{}, err
	}

	fromAddress, err := deriveFromAddress(privateKey)
	if err != nil {
		return nil, common.Address{}, err
	}

	nonce, err := s.ethClient.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, common.Address{}, err
	}
	gasPrice, err := s.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, common.Address{}, err
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

func (s *Shell) GetTxSender(ctx context.Context, txHash common.Hash) (common.Address, error) {
	tx, _, err := s.ethClient.TransactionByHash(ctx, txHash)
	if err != nil {
		return common.Address{}, err
	}

	chainID, err := s.ethClient.NetworkID(ctx)
	if err != nil {
		return common.Address{}, err
	}

	msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil)
	if err != nil {
		return common.Address{}, err
	}

	return msg.From(), nil
}

func (s *Shell) HeaderByNumber(ctx context.Context, blockNumber uint64) (*types.Header, error) {
	header, err := s.ethClient.HeaderByNumber(ctx, big.NewInt(int64(blockNumber)))
	return header, err
}
