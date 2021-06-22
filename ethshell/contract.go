// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethshell

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EventLogEvent is an auto generated low-level Go binding around an user-defined struct.
type EventLogEvent struct {
	Cid         string
	NewAccounts []common.Address
	Emitter     common.Address
	Timestamp   *big.Int
	BlockNumber *big.Int
}

// DelABI is the input ABI used to generate the binding from.
const DelABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"EventPublished\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"}],\"name\":\"getEvents\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structEventLog.Event[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTopics\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"publishEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DelBin is the compiled bytecode used for deploying new contracts.
var DelBin = "0x608060405234801561001057600080fd5b506117d7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063572bfd9614610046578063e18f882114610062578063f80e1a3c14610092575b600080fd5b610060600480360381019061005b9190610ec6565b6100c2565b005b61007c60048036038101906100779190610e5c565b61014f565b60405161008991906112de565b60405180910390f35b6100ac60048036038101906100a79190610e85565b610267565b6040516100b99190611300565b60405180910390f35b6000806000856040516100d591906112c7565b90815260200160405180910390206000018054905011905080610102576100fd848484610471565b61010e565b61010d8484846107fc565b5b7faa562e0c991226d0dd4de2957b8e3a190998a878210e2143466ae8123d96195784848460405161014193929190611322565b60405180910390a150505050565b6060600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561025c5783829060005260206000200180546101cf906115a2565b80601f01602080910402602001604051908101604052809291908181526020018280546101fb906115a2565b80156102485780601f1061021d57610100808354040283529160200191610248565b820191906000526020600020905b81548152906001019060200180831161022b57829003601f168201915b5050505050815260200190600101906101b0565b505050509050919050565b606060008260405161027991906112c7565b9081526020016040518091039020600001805480602002602001604051908101604052809291908181526020016000905b8282101561046657838290600052602060002090600502016040518060a00160405290816000820180546102dd906115a2565b80601f0160208091040260200160405190810160405280929190818152602001828054610309906115a2565b80156103565780601f1061032b57610100808354040283529160200191610356565b820191906000526020600020905b81548152906001019060200180831161033957829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156103e457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161039a575b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481525050815260200190600101906102aa565b505050509050919050565b6000808460405161048291906112c7565b908152602001604051809103902090506000846040516104a291906112c7565b90815260200160405180910390206000016040518060a001604052808581526020018481526020013273ffffffffffffffffffffffffffffffffffffffff1681526020014281526020014381525090806001815401808255809150506001900390600052602060002090600502016000909190919091506000820151816000019080519060200190610535929190610c1c565b506020820151816001019080519060200190610552929190610ca2565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015560808201518160040155505060018160010160003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610682929190610c1c565b5060005b82518110156107f55760018260010160008584815181106106d0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060016000848381518110610766577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020859080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906107e1929190610c1c565b5080806107ed90611605565b915050610686565b5050505050565b6108068332610ba7565b610845576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083c9061136e565b60405180910390fd5b60005b81518110156108ee5761089b8483838151811061088e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151610ba7565b156108db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d29061138e565b60405180910390fd5b80806108e690611605565b915050610848565b506000808460405161090091906112c7565b9081526020016040518091039020905060008460405161092091906112c7565b90815260200160405180910390206000016040518060a001604052808581526020018481526020013273ffffffffffffffffffffffffffffffffffffffff16815260200142815260200143815250908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000190805190602001906109b3929190610c1c565b5060208201518160010190805190602001906109d0929190610ca2565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015560808201518160040155505060005b8251811015610ba0576001826001016000858481518110610a7b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060016000848381518110610b11577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002085908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610b8c929190610c1c565b508080610b9890611605565b915050610a31565b5050505050565b60008083604051610bb891906112c7565b908152602001604051809103902060010160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b828054610c28906115a2565b90600052602060002090601f016020900481019282610c4a5760008555610c91565b82601f10610c6357805160ff1916838001178555610c91565b82800160010185558215610c91579182015b82811115610c90578251825591602001919060010190610c75565b5b509050610c9e9190610d2c565b5090565b828054828255906000526020600020908101928215610d1b579160200282015b82811115610d1a5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610cc2565b5b509050610d289190610d2c565b5090565b5b80821115610d45576000816000905550600101610d2d565b5090565b6000610d5c610d57846113d3565b6113ae565b90508083825260208201905082856020860282011115610d7b57600080fd5b60005b85811015610dab5781610d918882610df3565b845260208401935060208301925050600181019050610d7e565b5050509392505050565b6000610dc8610dc3846113ff565b6113ae565b905082815260208101848484011115610de057600080fd5b610deb848285611560565b509392505050565b600081359050610e028161178a565b92915050565b600082601f830112610e1957600080fd5b8135610e29848260208601610d49565b91505092915050565b600082601f830112610e4357600080fd5b8135610e53848260208601610db5565b91505092915050565b600060208284031215610e6e57600080fd5b6000610e7c84828501610df3565b91505092915050565b600060208284031215610e9757600080fd5b600082013567ffffffffffffffff811115610eb157600080fd5b610ebd84828501610e32565b91505092915050565b600080600060608486031215610edb57600080fd5b600084013567ffffffffffffffff811115610ef557600080fd5b610f0186828701610e32565b935050602084013567ffffffffffffffff811115610f1e57600080fd5b610f2a86828701610e32565b925050604084013567ffffffffffffffff811115610f4757600080fd5b610f5386828701610e08565b9150509250925092565b6000610f698383610f9d565b60208301905092915050565b6000610f818383611152565b905092915050565b6000610f95838361123b565b905092915050565b610fa681611524565b82525050565b6000610fb782611460565b610fc181856114b3565b9350610fcc83611430565b8060005b83811015610ffd578151610fe48882610f5d565b9750610fef8361148c565b925050600181019050610fd0565b5085935050505092915050565b600061101582611460565b61101f81856114c4565b935061102a83611430565b8060005b8381101561105b5781516110428882610f5d565b975061104d8361148c565b92505060018101905061102e565b5085935050505092915050565b60006110738261146b565b61107d81856114d5565b93508360208202850161108f85611440565b8060005b858110156110cb57848403895281516110ac8582610f75565b94506110b783611499565b925060208a01995050600181019050611093565b50829750879550505050505092915050565b60006110e882611476565b6110f281856114e6565b93508360208202850161110485611450565b8060005b8581101561114057848403895281516111218582610f89565b945061112c836114a6565b925060208a01995050600181019050611108565b50829750879550505050505092915050565b600061115d82611481565b61116781856114f7565b935061117781856020860161156f565b611180816116db565b840191505092915050565b600061119682611481565b6111a08185611508565b93506111b081856020860161156f565b6111b9816116db565b840191505092915050565b60006111cf82611481565b6111d98185611519565b93506111e981856020860161156f565b80840191505092915050565b6000611202602483611508565b915061120d826116ec565b604082019050919050565b6000611225602983611508565b91506112308261173b565b604082019050919050565b600060a08301600083015184820360008601526112588282611152565b915050602083015184820360208601526112728282610fac565b91505060408301516112876040860182610f9d565b50606083015161129a60608601826112b8565b5060808301516112ad60808601826112b8565b508091505092915050565b6112c181611556565b82525050565b60006112d382846111c4565b915081905092915050565b600060208201905081810360008301526112f88184611068565b905092915050565b6000602082019050818103600083015261131a81846110dd565b905092915050565b6000606082019050818103600083015261133c818661118b565b90508181036020830152611350818561118b565b90508181036040830152611364818461100a565b9050949350505050565b60006020820190508181036000830152611387816111f5565b9050919050565b600060208201905081810360008301526113a781611218565b9050919050565b60006113b86113c9565b90506113c482826115d4565b919050565b6000604051905090565b600067ffffffffffffffff8211156113ee576113ed6116ac565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561141a576114196116ac565b5b611423826116db565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061152f82611536565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561158d578082015181840152602081019050611572565b8381111561159c576000848401525b50505050565b600060028204905060018216806115ba57607f821691505b602082108114156115ce576115cd61167d565b5b50919050565b6115dd826116db565b810181811067ffffffffffffffff821117156115fc576115fb6116ac565b5b80604052505050565b600061161082611556565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156116435761164261164e565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f53656e646572206973206e6f742061206d656d626572206f662074686973207460008201527f6f70696300000000000000000000000000000000000000000000000000000000602082015250565b7f4163636f756e7420697320616c72656164792061206d656d626572206f66207460008201527f68697320746f7069630000000000000000000000000000000000000000000000602082015250565b61179381611524565b811461179e57600080fd5b5056fea2646970667358221220ad0a5558dd6c777353eea106087e42a655a52e3cd1f5d395905518bbe8c03cdc64736f6c63430008020033"

// DeployDel deploys a new Ethereum contract, binding an instance of Del to it.
func DeployDel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Del, error) {
	parsed, err := abi.JSON(strings.NewReader(DelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Del{DelCaller: DelCaller{contract: contract}, DelTransactor: DelTransactor{contract: contract}, DelFilterer: DelFilterer{contract: contract}}, nil
}

// Del is an auto generated Go binding around an Ethereum contract.
type Del struct {
	DelCaller     // Read-only binding to the contract
	DelTransactor // Write-only binding to the contract
	DelFilterer   // Log filterer for contract events
}

// DelCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelSession struct {
	Contract     *Del              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelCallerSession struct {
	Contract *DelCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelTransactorSession struct {
	Contract     *DelTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelRaw struct {
	Contract *Del // Generic contract binding to access the raw methods on
}

// DelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelCallerRaw struct {
	Contract *DelCaller // Generic read-only contract binding to access the raw methods on
}

// DelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelTransactorRaw struct {
	Contract *DelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDel creates a new instance of Del, bound to a specific deployed contract.
func NewDel(address common.Address, backend bind.ContractBackend) (*Del, error) {
	contract, err := bindDel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Del{DelCaller: DelCaller{contract: contract}, DelTransactor: DelTransactor{contract: contract}, DelFilterer: DelFilterer{contract: contract}}, nil
}

// NewDelCaller creates a new read-only instance of Del, bound to a specific deployed contract.
func NewDelCaller(address common.Address, caller bind.ContractCaller) (*DelCaller, error) {
	contract, err := bindDel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelCaller{contract: contract}, nil
}

// NewDelTransactor creates a new write-only instance of Del, bound to a specific deployed contract.
func NewDelTransactor(address common.Address, transactor bind.ContractTransactor) (*DelTransactor, error) {
	contract, err := bindDel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelTransactor{contract: contract}, nil
}

// NewDelFilterer creates a new log filterer instance of Del, bound to a specific deployed contract.
func NewDelFilterer(address common.Address, filterer bind.ContractFilterer) (*DelFilterer, error) {
	contract, err := bindDel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelFilterer{contract: contract}, nil
}

// bindDel binds a generic wrapper to an already deployed contract.
func bindDel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Del *DelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Del.Contract.DelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Del *DelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Del.Contract.DelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Del *DelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Del.Contract.DelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Del *DelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Del.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Del *DelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Del.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Del *DelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Del.Contract.contract.Transact(opts, method, params...)
}

// GetEvents is a free data retrieval call binding the contract method 0xf80e1a3c.
//
// Solidity: function getEvents(string topicId) view returns((string,address[],address,uint256,uint256)[])
func (_Del *DelCaller) GetEvents(opts *bind.CallOpts, topicId string) ([]EventLogEvent, error) {
	var out []interface{}
	err := _Del.contract.Call(opts, &out, "getEvents", topicId)

	if err != nil {
		return *new([]EventLogEvent), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventLogEvent)).(*[]EventLogEvent)

	return out0, err

}

// GetEvents is a free data retrieval call binding the contract method 0xf80e1a3c.
//
// Solidity: function getEvents(string topicId) view returns((string,address[],address,uint256,uint256)[])
func (_Del *DelSession) GetEvents(topicId string) ([]EventLogEvent, error) {
	return _Del.Contract.GetEvents(&_Del.CallOpts, topicId)
}

// GetEvents is a free data retrieval call binding the contract method 0xf80e1a3c.
//
// Solidity: function getEvents(string topicId) view returns((string,address[],address,uint256,uint256)[])
func (_Del *DelCallerSession) GetEvents(topicId string) ([]EventLogEvent, error) {
	return _Del.Contract.GetEvents(&_Del.CallOpts, topicId)
}

// GetTopics is a free data retrieval call binding the contract method 0xe18f8821.
//
// Solidity: function getTopics(address account) view returns(string[])
func (_Del *DelCaller) GetTopics(opts *bind.CallOpts, account common.Address) ([]string, error) {
	var out []interface{}
	err := _Del.contract.Call(opts, &out, "getTopics", account)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetTopics is a free data retrieval call binding the contract method 0xe18f8821.
//
// Solidity: function getTopics(address account) view returns(string[])
func (_Del *DelSession) GetTopics(account common.Address) ([]string, error) {
	return _Del.Contract.GetTopics(&_Del.CallOpts, account)
}

// GetTopics is a free data retrieval call binding the contract method 0xe18f8821.
//
// Solidity: function getTopics(address account) view returns(string[])
func (_Del *DelCallerSession) GetTopics(account common.Address) ([]string, error) {
	return _Del.Contract.GetTopics(&_Del.CallOpts, account)
}

// PublishEvent is a paid mutator transaction binding the contract method 0x572bfd96.
//
// Solidity: function publishEvent(string topicId, string cid, address[] newAccounts) returns()
func (_Del *DelTransactor) PublishEvent(opts *bind.TransactOpts, topicId string, cid string, newAccounts []common.Address) (*types.Transaction, error) {
	return _Del.contract.Transact(opts, "publishEvent", topicId, cid, newAccounts)
}

// PublishEvent is a paid mutator transaction binding the contract method 0x572bfd96.
//
// Solidity: function publishEvent(string topicId, string cid, address[] newAccounts) returns()
func (_Del *DelSession) PublishEvent(topicId string, cid string, newAccounts []common.Address) (*types.Transaction, error) {
	return _Del.Contract.PublishEvent(&_Del.TransactOpts, topicId, cid, newAccounts)
}

// PublishEvent is a paid mutator transaction binding the contract method 0x572bfd96.
//
// Solidity: function publishEvent(string topicId, string cid, address[] newAccounts) returns()
func (_Del *DelTransactorSession) PublishEvent(topicId string, cid string, newAccounts []common.Address) (*types.Transaction, error) {
	return _Del.Contract.PublishEvent(&_Del.TransactOpts, topicId, cid, newAccounts)
}

// DelEventPublishedIterator is returned from FilterEventPublished and is used to iterate over the raw logs and unpacked data for EventPublished events raised by the Del contract.
type DelEventPublishedIterator struct {
	Event *DelEventPublished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DelEventPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelEventPublished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DelEventPublished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DelEventPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelEventPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelEventPublished represents a EventPublished event raised by the Del contract.
type DelEventPublished struct {
	TopicId     string
	Cid         string
	NewAccounts []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEventPublished is a free log retrieval operation binding the contract event 0xaa562e0c991226d0dd4de2957b8e3a190998a878210e2143466ae8123d961957.
//
// Solidity: event EventPublished(string topicId, string cid, address[] newAccounts)
func (_Del *DelFilterer) FilterEventPublished(opts *bind.FilterOpts) (*DelEventPublishedIterator, error) {

	logs, sub, err := _Del.contract.FilterLogs(opts, "EventPublished")
	if err != nil {
		return nil, err
	}
	return &DelEventPublishedIterator{contract: _Del.contract, event: "EventPublished", logs: logs, sub: sub}, nil
}

// WatchEventPublished is a free log subscription operation binding the contract event 0xaa562e0c991226d0dd4de2957b8e3a190998a878210e2143466ae8123d961957.
//
// Solidity: event EventPublished(string topicId, string cid, address[] newAccounts)
func (_Del *DelFilterer) WatchEventPublished(opts *bind.WatchOpts, sink chan<- *DelEventPublished) (event.Subscription, error) {

	logs, sub, err := _Del.contract.WatchLogs(opts, "EventPublished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelEventPublished)
				if err := _Del.contract.UnpackLog(event, "EventPublished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventPublished is a log parse operation binding the contract event 0xaa562e0c991226d0dd4de2957b8e3a190998a878210e2143466ae8123d961957.
//
// Solidity: event EventPublished(string topicId, string cid, address[] newAccounts)
func (_Del *DelFilterer) ParseEventPublished(log types.Log) (*DelEventPublished, error) {
	event := new(DelEventPublished)
	if err := _Del.contract.UnpackLog(event, "EventPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
