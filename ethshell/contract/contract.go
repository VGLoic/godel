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
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"EventPublished\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"}],\"name\":\"getEvents\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structEventLog.Event[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTopics\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"publishEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x608060405234801561001057600080fd5b5061179d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063572bfd9614610046578063e18f882114610062578063f80e1a3c14610092575b600080fd5b610060600480360381019061005b9190610e8c565b6100c2565b005b61007c60048036038101906100779190610e22565b61014f565b60405161008991906112a4565b60405180910390f35b6100ac60048036038101906100a79190610e4b565b610267565b6040516100b991906112c6565b60405180910390f35b6000806000856040516100d5919061128d565b90815260200160405180910390206000018054905011905080610102576100fd848484610471565b61010e565b61010d8484846107df565b5b7faa562e0c991226d0dd4de2957b8e3a190998a878210e2143466ae8123d961957848484604051610141939291906112e8565b60405180910390a150505050565b6060600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561025c5783829060005260206000200180546101cf90611568565b80601f01602080910402602001604051908101604052809291908181526020018280546101fb90611568565b80156102485780601f1061021d57610100808354040283529160200191610248565b820191906000526020600020905b81548152906001019060200180831161022b57829003601f168201915b5050505050815260200190600101906101b0565b505050509050919050565b6060600082604051610279919061128d565b9081526020016040518091039020600001805480602002602001604051908101604052809291908181526020016000905b8282101561046657838290600052602060002090600502016040518060a00160405290816000820180546102dd90611568565b80601f016020809104026020016040519081016040528092919081815260200182805461030990611568565b80156103565780601f1061032b57610100808354040283529160200191610356565b820191906000526020600020905b81548152906001019060200180831161033957829003601f168201915b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156103e457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161039a575b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481525050815260200190600101906102aa565b505050509050919050565b60008084604051610482919061128d565b90815260200160405180910390209050806000016040518060a001604052808581526020018481526020013273ffffffffffffffffffffffffffffffffffffffff1681526020014281526020014381525090806001815401808255809150506001900390600052602060002090600502016000909190919091506000820151816000019080519060200190610518929190610be2565b506020820151816001019080519060200190610535929190610c68565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015560808201518160040155505060018160010160003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610665929190610be2565b5060005b82518110156107d85760018260010160008584815181106106b3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060016000848381518110610749577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020859080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906107c4929190610be2565b5080806107d0906115cb565b915050610669565b5050505050565b6107e98332610b6d565b610828576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081f90611334565b60405180910390fd5b60005b81518110156108d15761087e84838381518110610871577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151610b6d565b156108be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b590611354565b60405180910390fd5b80806108c9906115cb565b91505061082b565b50600080846040516108e3919061128d565b90815260200160405180910390209050806000016040518060a001604052808581526020018481526020013273ffffffffffffffffffffffffffffffffffffffff1681526020014281526020014381525090806001815401808255809150506001900390600052602060002090600502016000909190919091506000820151816000019080519060200190610979929190610be2565b506020820151816001019080519060200190610996929190610c68565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003015560808201518160040155505060005b8251811015610b66576001826001016000858481518110610a41577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060016000848381518110610ad7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002085908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610b52929190610be2565b508080610b5e906115cb565b9150506109f7565b5050505050565b60008083604051610b7e919061128d565b908152602001604051809103902060010160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b828054610bee90611568565b90600052602060002090601f016020900481019282610c105760008555610c57565b82601f10610c2957805160ff1916838001178555610c57565b82800160010185558215610c57579182015b82811115610c56578251825591602001919060010190610c3b565b5b509050610c649190610cf2565b5090565b828054828255906000526020600020908101928215610ce1579160200282015b82811115610ce05782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190610c88565b5b509050610cee9190610cf2565b5090565b5b80821115610d0b576000816000905550600101610cf3565b5090565b6000610d22610d1d84611399565b611374565b90508083825260208201905082856020860282011115610d4157600080fd5b60005b85811015610d715781610d578882610db9565b845260208401935060208301925050600181019050610d44565b5050509392505050565b6000610d8e610d89846113c5565b611374565b905082815260208101848484011115610da657600080fd5b610db1848285611526565b509392505050565b600081359050610dc881611750565b92915050565b600082601f830112610ddf57600080fd5b8135610def848260208601610d0f565b91505092915050565b600082601f830112610e0957600080fd5b8135610e19848260208601610d7b565b91505092915050565b600060208284031215610e3457600080fd5b6000610e4284828501610db9565b91505092915050565b600060208284031215610e5d57600080fd5b600082013567ffffffffffffffff811115610e7757600080fd5b610e8384828501610df8565b91505092915050565b600080600060608486031215610ea157600080fd5b600084013567ffffffffffffffff811115610ebb57600080fd5b610ec786828701610df8565b935050602084013567ffffffffffffffff811115610ee457600080fd5b610ef086828701610df8565b925050604084013567ffffffffffffffff811115610f0d57600080fd5b610f1986828701610dce565b9150509250925092565b6000610f2f8383610f63565b60208301905092915050565b6000610f478383611118565b905092915050565b6000610f5b8383611201565b905092915050565b610f6c816114ea565b82525050565b6000610f7d82611426565b610f878185611479565b9350610f92836113f6565b8060005b83811015610fc3578151610faa8882610f23565b9750610fb583611452565b925050600181019050610f96565b5085935050505092915050565b6000610fdb82611426565b610fe5818561148a565b9350610ff0836113f6565b8060005b838110156110215781516110088882610f23565b975061101383611452565b925050600181019050610ff4565b5085935050505092915050565b600061103982611431565b611043818561149b565b93508360208202850161105585611406565b8060005b8581101561109157848403895281516110728582610f3b565b945061107d8361145f565b925060208a01995050600181019050611059565b50829750879550505050505092915050565b60006110ae8261143c565b6110b881856114ac565b9350836020820285016110ca85611416565b8060005b8581101561110657848403895281516110e78582610f4f565b94506110f28361146c565b925060208a019950506001810190506110ce565b50829750879550505050505092915050565b600061112382611447565b61112d81856114bd565b935061113d818560208601611535565b611146816116a1565b840191505092915050565b600061115c82611447565b61116681856114ce565b9350611176818560208601611535565b61117f816116a1565b840191505092915050565b600061119582611447565b61119f81856114df565b93506111af818560208601611535565b80840191505092915050565b60006111c86024836114ce565b91506111d3826116b2565b604082019050919050565b60006111eb6029836114ce565b91506111f682611701565b604082019050919050565b600060a083016000830151848203600086015261121e8282611118565b915050602083015184820360208601526112388282610f72565b915050604083015161124d6040860182610f63565b506060830151611260606086018261127e565b506080830151611273608086018261127e565b508091505092915050565b6112878161151c565b82525050565b6000611299828461118a565b915081905092915050565b600060208201905081810360008301526112be818461102e565b905092915050565b600060208201905081810360008301526112e081846110a3565b905092915050565b600060608201905081810360008301526113028186611151565b905081810360208301526113168185611151565b9050818103604083015261132a8184610fd0565b9050949350505050565b6000602082019050818103600083015261134d816111bb565b9050919050565b6000602082019050818103600083015261136d816111de565b9050919050565b600061137e61138f565b905061138a828261159a565b919050565b6000604051905090565b600067ffffffffffffffff8211156113b4576113b3611672565b5b602082029050602081019050919050565b600067ffffffffffffffff8211156113e0576113df611672565b5b6113e9826116a1565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60006114f5826114fc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611553578082015181840152602081019050611538565b83811115611562576000848401525b50505050565b6000600282049050600182168061158057607f821691505b6020821081141561159457611593611643565b5b50919050565b6115a3826116a1565b810181811067ffffffffffffffff821117156115c2576115c1611672565b5b80604052505050565b60006115d68261151c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561160957611608611614565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f53656e646572206973206e6f742061206d656d626572206f662074686973207460008201527f6f70696300000000000000000000000000000000000000000000000000000000602082015250565b7f4163636f756e7420697320616c72656164792061206d656d626572206f66207460008201527f68697320746f7069630000000000000000000000000000000000000000000000602082015250565b611759816114ea565b811461176457600080fd5b5056fea2646970667358221220188868d601e3e45595eb0688a989f9876c96c22d1292eab53817a4bab8bfc8c964736f6c63430008020033"

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
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEventPublished is a free log retrieval operation binding the contract event 0xaa562e0c991226d0dd4de2957b8e3a190998a878210e2143466ae8123d961957.
//
// Solidity: event EventPublished(string topicId, string cid, address[] newAccounts)
func (_Contract *ContractFilterer) FilterEventPublished(opts *bind.FilterOpts) (*ContractEventPublishedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EventPublished")
	if err != nil {
		return nil, err
	}
	return &ContractEventPublishedIterator{contract: _Contract.contract, event: "EventPublished", logs: logs, sub: sub}, nil
}

// WatchEventPublished is a free log subscription operation binding the contract event 0xaa562e0c991226d0dd4de2957b8e3a190998a878210e2143466ae8123d961957.
//
// Solidity: event EventPublished(string topicId, string cid, address[] newAccounts)
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

// ParseEventPublished is a log parse operation binding the contract event 0xaa562e0c991226d0dd4de2957b8e3a190998a878210e2143466ae8123d961957.
//
// Solidity: event EventPublished(string topicId, string cid, address[] newAccounts)
func (_Contract *ContractFilterer) ParseEventPublished(log types.Log) (*ContractEventPublished, error) {
	event := new(ContractEventPublished)
	if err := _Contract.contract.UnpackLog(event, "EventPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
