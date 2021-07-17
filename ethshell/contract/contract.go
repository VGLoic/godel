// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"}],\"name\":\"EventPublished\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"}],\"name\":\"getEvents\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structEventLog.Event[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTopics\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"publishEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x608060405234801561001057600080fd5b5061157c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063572bfd9614610046578063e18f882114610062578063f80e1a3c14610092575b600080fd5b610060600480360381019061005b9190610c4e565b6100c2565b005b61007c60048036038101906100779190610be4565b61015a565b6040516100899190611075565b60405180910390f35b6100ac60048036038101906100a79190610c0d565b610272565b6040516100b99190611097565b60405180910390f35b6000806000856040516100d5919061105e565b908152602001604051809103902060000180549050119050600081610108576100ff85858561047c565b60019050610116565b61011385858561057d565b90505b7fd67081c19db8c97a6739b647c4a584692cd5a1376518bcd13bd789d2f10bbab38585858460405161014b94939291906110b9565b60405180910390a15050505050565b6060600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156102675783829060005260206000200180546101da90611347565b80601f016020809104026020016040519081016040528092919081815260200182805461020690611347565b80156102535780601f1061022857610100808354040283529160200191610253565b820191906000526020600020905b81548152906001019060200180831161023657829003601f168201915b5050505050815260200190600101906101bb565b505050509050919050565b6060600082604051610284919061105e565b9081526020016040518091039020600001805480602002602001604051908101604052809291908181526020016000905b8282101561047157838290600052602060002090600502016040518060a00160405290816000820180546102e890611347565b80601f016020809104026020016040519081016040528092919081815260200182805461031490611347565b80156103615780601f1061033657610100808354040283529160200191610361565b820191906000526020600020905b81548152906001019060200180831161034457829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156103ef57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116103a5575b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481525050815260200190600101906102b5565b505050509050919050565b600160008460405161048e919061105e565b908152602001604051809103902060010160003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208390806001815401808255809150506001900390600052602060002001600090919091909150908051906020019061056b9291906109a4565b50610577838383610686565b50505050565b6000610589843261092f565b6105c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105bf90611113565b60405180910390fd5b60005b82518110156106715761061e85848381518110610611577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015161092f565b1561065e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065590611133565b60405180910390fd5b8080610669906113aa565b9150506105cb565b5061067d848484610686565b90509392505050565b600080600085604051610699919061105e565b90815260200160405180910390209050806000016040518060a001604052808681526020018581526020013273ffffffffffffffffffffffffffffffffffffffff168152602001428152602001438152509080600181540180825580915050600190039060005260206000209060050201600090919091909150600082015181600001908051906020019061072f9291906109a4565b50602082015181600101908051906020019061074c929190610a2a565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015560808201518160040155505060005b835181101561091c5760018260010160008684815181106107f7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600085838151811061088d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020869080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906109089291906109a4565b508080610914906113aa565b9150506107ad565b5080600001805490509150509392505050565b60008083604051610940919061105e565b908152602001604051809103902060010160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b8280546109b090611347565b90600052602060002090601f0160209004810192826109d25760008555610a19565b82601f106109eb57805160ff1916838001178555610a19565b82800160010185558215610a19579182015b82811115610a185782518255916020019190600101906109fd565b5b509050610a269190610ab4565b5090565b828054828255906000526020600020908101928215610aa3579160200282015b82811115610aa25782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610a4a565b5b509050610ab09190610ab4565b5090565b5b80821115610acd576000816000905550600101610ab5565b5090565b6000610ae4610adf84611178565b611153565b90508083825260208201905082856020860282011115610b0357600080fd5b60005b85811015610b335781610b198882610b7b565b845260208401935060208301925050600181019050610b06565b5050509392505050565b6000610b50610b4b846111a4565b611153565b905082815260208101848484011115610b6857600080fd5b610b73848285611305565b509392505050565b600081359050610b8a8161152f565b92915050565b600082601f830112610ba157600080fd5b8135610bb1848260208601610ad1565b91505092915050565b600082601f830112610bcb57600080fd5b8135610bdb848260208601610b3d565b91505092915050565b600060208284031215610bf657600080fd5b6000610c0484828501610b7b565b91505092915050565b600060208284031215610c1f57600080fd5b600082013567ffffffffffffffff811115610c3957600080fd5b610c4584828501610bba565b91505092915050565b600080600060608486031215610c6357600080fd5b600084013567ffffffffffffffff811115610c7d57600080fd5b610c8986828701610bba565b935050602084013567ffffffffffffffff811115610ca657600080fd5b610cb286828701610bba565b925050604084013567ffffffffffffffff811115610ccf57600080fd5b610cdb86828701610b90565b9150509250925092565b6000610cf18383610d25565b60208301905092915050565b6000610d098383610eda565b905092915050565b6000610d1d8383610fc3565b905092915050565b610d2e816112c9565b82525050565b6000610d3f82611205565b610d498185611258565b9350610d54836111d5565b8060005b83811015610d85578151610d6c8882610ce5565b9750610d7783611231565b925050600181019050610d58565b5085935050505092915050565b6000610d9d82611205565b610da78185611269565b9350610db2836111d5565b8060005b83811015610de3578151610dca8882610ce5565b9750610dd583611231565b925050600181019050610db6565b5085935050505092915050565b6000610dfb82611210565b610e05818561127a565b935083602082028501610e17856111e5565b8060005b85811015610e535784840389528151610e348582610cfd565b9450610e3f8361123e565b925060208a01995050600181019050610e1b565b50829750879550505050505092915050565b6000610e708261121b565b610e7a818561128b565b935083602082028501610e8c856111f5565b8060005b85811015610ec85784840389528151610ea98582610d11565b9450610eb48361124b565b925060208a01995050600181019050610e90565b50829750879550505050505092915050565b6000610ee582611226565b610eef818561129c565b9350610eff818560208601611314565b610f0881611480565b840191505092915050565b6000610f1e82611226565b610f2881856112ad565b9350610f38818560208601611314565b610f4181611480565b840191505092915050565b6000610f5782611226565b610f6181856112be565b9350610f71818560208601611314565b80840191505092915050565b6000610f8a6024836112ad565b9150610f9582611491565b604082019050919050565b6000610fad6029836112ad565b9150610fb8826114e0565b604082019050919050565b600060a0830160008301518482036000860152610fe08282610eda565b91505060208301518482036020860152610ffa8282610d34565b915050604083015161100f6040860182610d25565b5060608301516110226060860182611040565b5060808301516110356080860182611040565b508091505092915050565b611049816112fb565b82525050565b611058816112fb565b82525050565b600061106a8284610f4c565b915081905092915050565b6000602082019050818103600083015261108f8184610df0565b905092915050565b600060208201905081810360008301526110b18184610e65565b905092915050565b600060808201905081810360008301526110d38187610f13565b905081810360208301526110e78186610f13565b905081810360408301526110fb8185610d92565b905061110a606083018461104f565b95945050505050565b6000602082019050818103600083015261112c81610f7d565b9050919050565b6000602082019050818103600083015261114c81610fa0565b9050919050565b600061115d61116e565b90506111698282611379565b919050565b6000604051905090565b600067ffffffffffffffff82111561119357611192611451565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156111bf576111be611451565b5b6111c882611480565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006112d4826112db565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611332578082015181840152602081019050611317565b83811115611341576000848401525b50505050565b6000600282049050600182168061135f57607f821691505b6020821081141561137357611372611422565b5b50919050565b61138282611480565b810181811067ffffffffffffffff821117156113a1576113a0611451565b5b80604052505050565b60006113b5826112fb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156113e8576113e76113f3565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f53656e646572206973206e6f742061206d656d626572206f662074686973207460008201527f6f70696300000000000000000000000000000000000000000000000000000000602082015250565b7f4163636f756e7420697320616c72656164792061206d656d626572206f66207460008201527f68697320746f7069630000000000000000000000000000000000000000000000602082015250565b611538816112c9565b811461154357600080fd5b5056fea26469706673582212208c39ca56a4f1b05e93dffc8e98b3dce8dba075023a861f9765ad54d960f1988764736f6c63430008020033"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetEvents is a free data retrieval call binding the contract method 0xf80e1a3c.
//
// Solidity: function getEvents(string topicId) view returns((string,address[],address,uint256,uint256)[])
func (_Contract *ContractCaller) GetEvents(opts *bind.CallOpts, topicId string) ([]EventLogEvent, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEvents", topicId)

	if err != nil {
		return *new([]EventLogEvent), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventLogEvent)).(*[]EventLogEvent)

	return out0, err

}

// GetEvents is a free data retrieval call binding the contract method 0xf80e1a3c.
//
// Solidity: function getEvents(string topicId) view returns((string,address[],address,uint256,uint256)[])
func (_Contract *ContractSession) GetEvents(topicId string) ([]EventLogEvent, error) {
	return _Contract.Contract.GetEvents(&_Contract.CallOpts, topicId)
}

// GetEvents is a free data retrieval call binding the contract method 0xf80e1a3c.
//
// Solidity: function getEvents(string topicId) view returns((string,address[],address,uint256,uint256)[])
func (_Contract *ContractCallerSession) GetEvents(topicId string) ([]EventLogEvent, error) {
	return _Contract.Contract.GetEvents(&_Contract.CallOpts, topicId)
}

// GetTopics is a free data retrieval call binding the contract method 0xe18f8821.
//
// Solidity: function getTopics(address account) view returns(string[])
func (_Contract *ContractCaller) GetTopics(opts *bind.CallOpts, account common.Address) ([]string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTopics", account)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetTopics is a free data retrieval call binding the contract method 0xe18f8821.
//
// Solidity: function getTopics(address account) view returns(string[])
func (_Contract *ContractSession) GetTopics(account common.Address) ([]string, error) {
	return _Contract.Contract.GetTopics(&_Contract.CallOpts, account)
}

// GetTopics is a free data retrieval call binding the contract method 0xe18f8821.
//
// Solidity: function getTopics(address account) view returns(string[])
func (_Contract *ContractCallerSession) GetTopics(account common.Address) ([]string, error) {
	return _Contract.Contract.GetTopics(&_Contract.CallOpts, account)
}

// PublishEvent is a paid mutator transaction binding the contract method 0x572bfd96.
//
// Solidity: function publishEvent(string topicId, string cid, address[] newAccounts) returns()
func (_Contract *ContractTransactor) PublishEvent(opts *bind.TransactOpts, topicId string, cid string, newAccounts []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "publishEvent", topicId, cid, newAccounts)
}

// PublishEvent is a paid mutator transaction binding the contract method 0x572bfd96.
//
// Solidity: function publishEvent(string topicId, string cid, address[] newAccounts) returns()
func (_Contract *ContractSession) PublishEvent(topicId string, cid string, newAccounts []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.PublishEvent(&_Contract.TransactOpts, topicId, cid, newAccounts)
}

// PublishEvent is a paid mutator transaction binding the contract method 0x572bfd96.
//
// Solidity: function publishEvent(string topicId, string cid, address[] newAccounts) returns()
func (_Contract *ContractTransactorSession) PublishEvent(topicId string, cid string, newAccounts []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.PublishEvent(&_Contract.TransactOpts, topicId, cid, newAccounts)
}

// ContractEventPublishedIterator is returned from FilterEventPublished and is used to iterate over the raw logs and unpacked data for EventPublished events raised by the Contract contract.
type ContractEventPublishedIterator struct {
	Event *ContractEventPublished // Event containing the contract specifics and raw log

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
func (it *ContractEventPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEventPublished)
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
		it.Event = new(ContractEventPublished)
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
func (it *ContractEventPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEventPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEventPublished represents a EventPublished event raised by the Contract contract.
type ContractEventPublished struct {
	TopicId     string
	Cid         string
	NewAccounts []common.Address
	Depth       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEventPublished is a free log retrieval operation binding the contract event 0xd67081c19db8c97a6739b647c4a584692cd5a1376518bcd13bd789d2f10bbab3.
//
// Solidity: event EventPublished(string topicId, string cid, address[] newAccounts, uint256 depth)
func (_Contract *ContractFilterer) FilterEventPublished(opts *bind.FilterOpts) (*ContractEventPublishedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EventPublished")
	if err != nil {
		return nil, err
	}
	return &ContractEventPublishedIterator{contract: _Contract.contract, event: "EventPublished", logs: logs, sub: sub}, nil
}

// WatchEventPublished is a free log subscription operation binding the contract event 0xd67081c19db8c97a6739b647c4a584692cd5a1376518bcd13bd789d2f10bbab3.
//
// Solidity: event EventPublished(string topicId, string cid, address[] newAccounts, uint256 depth)
func (_Contract *ContractFilterer) WatchEventPublished(opts *bind.WatchOpts, sink chan<- *ContractEventPublished) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EventPublished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEventPublished)
				if err := _Contract.contract.UnpackLog(event, "EventPublished", log); err != nil {
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

// ParseEventPublished is a log parse operation binding the contract event 0xd67081c19db8c97a6739b647c4a584692cd5a1376518bcd13bd789d2f10bbab3.
//
// Solidity: event EventPublished(string topicId, string cid, address[] newAccounts, uint256 depth)
func (_Contract *ContractFilterer) ParseEventPublished(log types.Log) (*ContractEventPublished, error) {
	event := new(ContractEventPublished)
	if err := _Contract.contract.UnpackLog(event, "EventPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
