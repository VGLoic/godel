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
	Id          []byte
	Next        []byte
}

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"EventPublished\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fromBlock\",\"type\":\"uint256\"}],\"name\":\"getEvents\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"emitter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"next\",\"type\":\"bytes\"}],\"internalType\":\"structEventLog.Event[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTopics\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"topicId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"newAccounts\",\"type\":\"address[]\"}],\"name\":\"publishEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x608060405234801561001057600080fd5b50612500806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806363ab919d14610046578063cc35570814610076578063e18f882114610092575b600080fd5b610060600480360381019061005b919061194d565b6100c2565b60405161006d9190611e78565b60405180910390f35b610090600480360381019061008b919061188a565b61098e565b005b6100ac60048036038101906100a79190611861565b610add565b6040516100b99190611e56565b60405180910390f35b6060600080846040516100d59190611e3f565b908152602001604051809103902060000180546100f190612262565b80601f016020809104026020016040519081016040528092919081815260200182805461011d90612262565b801561016a5780601f1061013f5761010080835404028352916020019161016a565b820191906000526020600020905b81548152906001019060200180831161014d57829003601f168201915b50505050509050600080856040516101829190611e3f565b908152602001604051809103902060010154905060008167ffffffffffffffff8111156101d8577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561021157816020015b6101fe6114cb565b8152602001906001900390816101f65790505b50905060006001846040516102269190611e28565b90815260200160405180910390206040518060e001604052908160008201805461024f90612262565b80601f016020809104026020016040519081016040528092919081815260200182805461027b90612262565b80156102c85780601f1061029d576101008083540402835291602001916102c8565b820191906000526020600020905b8154815290600101906020018083116102ab57829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561035657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161030c575b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160038201548152602001600482015481526020016005820180546103d990612262565b80601f016020809104026020016040519081016040528092919081815260200182805461040590612262565b80156104525780601f1061042757610100808354040283529160200191610452565b820191906000526020600020905b81548152906001019060200180831161043557829003601f168201915b5050505050815260200160068201805461046b90612262565b80601f016020809104026020016040519081016040528092919081815260200182805461049790612262565b80156104e45780601f106104b9576101008083540402835291602001916104e4565b820191906000526020600020905b8154815290600101906020018083116104c757829003601f168201915b505050505081525050905060005b60008260a001515111801561050b575086826080015110155b15610846578183828151811061054a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525060018260c001516040516105699190611e28565b90815260200160405180910390206040518060e001604052908160008201805461059290612262565b80601f01602080910402602001604051908101604052809291908181526020018280546105be90612262565b801561060b5780601f106105e05761010080835404028352916020019161060b565b820191906000526020600020905b8154815290600101906020018083116105ee57829003601f168201915b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561069957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161064f575b505050505081526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382015481526020016004820154815260200160058201805461071c90612262565b80601f016020809104026020016040519081016040528092919081815260200182805461074890612262565b80156107955780601f1061076a57610100808354040283529160200191610795565b820191906000526020600020905b81548152906001019060200180831161077857829003601f168201915b505050505081526020016006820180546107ae90612262565b80601f01602080910402602001604051908101604052809291908181526020018280546107da90612262565b80156108275780601f106107fc57610100808354040283529160200191610827565b820191906000526020600020905b81548152906001019060200180831161080a57829003601f168201915b505050505081525050915060018161083f919061215a565b90506104f2565b60008167ffffffffffffffff811115610888577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156108c157816020015b6108ae6114cb565b8152602001906001900390816108a65790505b50905060005b8281101561097e57848181518110610908577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151828260018661091f91906121b0565b61092991906121b0565b81518110610960577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101819052508080610976906122c5565b9150506108c7565b5080965050505050505092915050565b60008351116109d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c990611f5b565b60405180910390fd5b60006001846040516109e49190611e28565b90815260200160405180910390206005018054610a0090612262565b905014610a42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3990611f1b565b60405180910390fd5b600080600086604051610a559190611e3f565b90815260200160405180910390206000018054610a7190612262565b905011905080610a8c57610a8785858585610bf5565b610a99565b610a9885858585610fc8565b5b7fbc285462bdca2ad65ad23ab3343b809e92ca2051ce3b33440b13833147aab8e785858585604051610ace9493929190611e9a565b60405180910390a15050505050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015610bea578382906000526020600020018054610b5d90612262565b80601f0160208091040260200160405190810160405280929190818152602001828054610b8990612262565b8015610bd65780601f10610bab57610100808354040283529160200191610bd6565b820191906000526020600020905b815481529060010190602001808311610bb957829003601f168201915b505050505081526020019060010190610b3e565b505050509050919050565b60008085604051610c069190611e3f565b90815260200160405180910390209050606060006040518060e001604052808681526020018581526020013273ffffffffffffffffffffffffffffffffffffffff16815260200142815260200143815260200187815260200183815250905080600187604051610c769190611e28565b90815260200160405180910390206000820151816000019080519060200190610ca092919061151e565b506020820151816001019080519060200190610cbd9291906115a4565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005019080519060200190610d3592919061162e565b5060c0820151816006019080519060200190610d5292919061162e565b5090505085836000019080519060200190610d6e92919061162e565b506001836001018190555060018360020160003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600260003273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002087908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610e4b92919061151e565b5060005b8451811015610fbe576001846002016000878481518110610e99577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060026000868381518110610f2f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002088908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610faa92919061151e565b508080610fb6906122c5565b915050610e4f565b5050505050505050565b610fd28432611456565b611011576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161100890611efb565b60405180910390fd5b60005b81518110156110ba576110678583838151811061105a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151611456565b156110a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161109e90611f3b565b60405180910390fd5b80806110b2906122c5565b915050611014565b50600080856040516110cc9190611e3f565b9081526020016040518091039020905060006040518060e001604052808581526020018481526020013273ffffffffffffffffffffffffffffffffffffffff16815260200142815260200143815260200186815260200183600001805461113290612262565b80601f016020809104026020016040519081016040528092919081815260200182805461115e90612262565b80156111ab5780601f10611180576101008083540402835291602001916111ab565b820191906000526020600020905b81548152906001019060200180831161118e57829003601f168201915b50505050508152509050806001866040516111c69190611e28565b908152602001604051809103902060008201518160000190805190602001906111f092919061151e565b50602082015181600101908051906020019061120d9291906115a4565b5060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a082015181600501908051906020019061128592919061162e565b5060c08201518160060190805190602001906112a292919061162e565b50905050848260000190805190602001906112be92919061162e565b5060018260010160008282546112d4919061215a565b9250508190555060005b835181101561144d576001836002016000868481518110611328577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600260008583815181106113be577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208790806001815401808255809150506001900390600052602060002001600090919091909150908051906020019061143992919061151e565b508080611445906122c5565b9150506112de565b50505050505050565b600080836040516114679190611e3f565b908152602001604051809103902060020160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6040518060e001604052806060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815260200160608152602001606081525090565b82805461152a90612262565b90600052602060002090601f01602090048101928261154c5760008555611593565b82601f1061156557805160ff1916838001178555611593565b82800160010185558215611593579182015b82811115611592578251825591602001919060010190611577565b5b5090506115a091906116b4565b5090565b82805482825590600052602060002090810192821561161d579160200282015b8281111561161c5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906115c4565b5b50905061162a91906116b4565b5090565b82805461163a90612262565b90600052602060002090601f01602090048101928261165c57600085556116a3565b82601f1061167557805160ff19168380011785556116a3565b828001600101855582156116a3579182015b828111156116a2578251825591602001919060010190611687565b5b5090506116b091906116b4565b5090565b5b808211156116cd5760008160009055506001016116b5565b5090565b60006116e46116df84611fa0565b611f7b565b9050808382526020820190508285602086028201111561170357600080fd5b60005b85811015611733578161171988826117b9565b845260208401935060208301925050600181019050611706565b5050509392505050565b600061175061174b84611fcc565b611f7b565b90508281526020810184848401111561176857600080fd5b611773848285612220565b509392505050565b600061178e61178984611ffd565b611f7b565b9050828152602081018484840111156117a657600080fd5b6117b1848285612220565b509392505050565b6000813590506117c88161249c565b92915050565b600082601f8301126117df57600080fd5b81356117ef8482602086016116d1565b91505092915050565b600082601f83011261180957600080fd5b813561181984826020860161173d565b91505092915050565b600082601f83011261183357600080fd5b813561184384826020860161177b565b91505092915050565b60008135905061185b816124b3565b92915050565b60006020828403121561187357600080fd5b6000611881848285016117b9565b91505092915050565b600080600080608085870312156118a057600080fd5b600085013567ffffffffffffffff8111156118ba57600080fd5b6118c687828801611822565b945050602085013567ffffffffffffffff8111156118e357600080fd5b6118ef878288016117f8565b935050604085013567ffffffffffffffff81111561190c57600080fd5b61191887828801611822565b925050606085013567ffffffffffffffff81111561193557600080fd5b611941878288016117ce565b91505092959194509250565b6000806040838503121561196057600080fd5b600083013567ffffffffffffffff81111561197a57600080fd5b61198685828601611822565b92505060206119978582860161184c565b9150509250929050565b60006119ad83836119e1565b60208301905092915050565b60006119c58383611c39565b905092915050565b60006119d98383611d68565b905092915050565b6119ea816121e4565b82525050565b60006119fb8261205e565b611a0581856120bc565b9350611a108361202e565b8060005b83811015611a41578151611a2888826119a1565b9750611a3383612095565b925050600181019050611a14565b5085935050505092915050565b6000611a598261205e565b611a6381856120cd565b9350611a6e8361202e565b8060005b83811015611a9f578151611a8688826119a1565b9750611a9183612095565b925050600181019050611a72565b5085935050505092915050565b6000611ab782612069565b611ac181856120de565b935083602082028501611ad38561203e565b8060005b85811015611b0f5784840389528151611af085826119b9565b9450611afb836120a2565b925060208a01995050600181019050611ad7565b50829750879550505050505092915050565b6000611b2c82612074565b611b3681856120ef565b935083602082028501611b488561204e565b8060005b85811015611b845784840389528151611b6585826119cd565b9450611b70836120af565b925060208a01995050600181019050611b4c565b50829750879550505050505092915050565b6000611ba18261207f565b611bab8185612100565b9350611bbb81856020860161222f565b611bc48161239b565b840191505092915050565b6000611bda8261207f565b611be48185612111565b9350611bf481856020860161222f565b611bfd8161239b565b840191505092915050565b6000611c138261207f565b611c1d8185612122565b9350611c2d81856020860161222f565b80840191505092915050565b6000611c448261208a565b611c4e818561212d565b9350611c5e81856020860161222f565b611c678161239b565b840191505092915050565b6000611c7d8261208a565b611c87818561213e565b9350611c9781856020860161222f565b611ca08161239b565b840191505092915050565b6000611cb68261208a565b611cc0818561214f565b9350611cd081856020860161222f565b80840191505092915050565b6000611ce960248361213e565b9150611cf4826123ac565b604082019050919050565b6000611d0c60148361213e565b9150611d17826123fb565b602082019050919050565b6000611d2f60298361213e565b9150611d3a82612424565b604082019050919050565b6000611d52600b8361213e565b9150611d5d82612473565b602082019050919050565b600060e0830160008301518482036000860152611d858282611c39565b91505060208301518482036020860152611d9f82826119f0565b9150506040830151611db460408601826119e1565b506060830151611dc76060860182611e19565b506080830151611dda6080860182611e19565b5060a083015184820360a0860152611df28282611b96565b91505060c083015184820360c0860152611e0c8282611b96565b9150508091505092915050565b611e2281612216565b82525050565b6000611e348284611c08565b915081905092915050565b6000611e4b8284611cab565b915081905092915050565b60006020820190508181036000830152611e708184611aac565b905092915050565b60006020820190508181036000830152611e928184611b21565b905092915050565b60006080820190508181036000830152611eb48187611c72565b90508181036020830152611ec88186611bcf565b90508181036040830152611edc8185611c72565b90508181036060830152611ef08184611a4e565b905095945050505050565b60006020820190508181036000830152611f1481611cdc565b9050919050565b60006020820190508181036000830152611f3481611cff565b9050919050565b60006020820190508181036000830152611f5481611d22565b9050919050565b60006020820190508181036000830152611f7481611d45565b9050919050565b6000611f85611f96565b9050611f918282612294565b919050565b6000604051905090565b600067ffffffffffffffff821115611fbb57611fba61236c565b5b602082029050602081019050919050565b600067ffffffffffffffff821115611fe757611fe661236c565b5b611ff08261239b565b9050602081019050919050565b600067ffffffffffffffff8211156120185761201761236c565b5b6120218261239b565b9050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061216582612216565b915061217083612216565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156121a5576121a461230e565b5b828201905092915050565b60006121bb82612216565b91506121c683612216565b9250828210156121d9576121d861230e565b5b828203905092915050565b60006121ef826121f6565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561224d578082015181840152602081019050612232565b8381111561225c576000848401525b50505050565b6000600282049050600182168061227a57607f821691505b6020821081141561228e5761228d61233d565b5b50919050565b61229d8261239b565b810181811067ffffffffffffffff821117156122bc576122bb61236c565b5b80604052505050565b60006122d082612216565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156123035761230261230e565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f53656e646572206973206e6f742061206d656d626572206f662074686973207460008201527f6f70696300000000000000000000000000000000000000000000000000000000602082015250565b7f4576656e7420616c726561647920657869737473000000000000000000000000600082015250565b7f4163636f756e7420697320616c72656164792061206d656d626572206f66207460008201527f68697320746f7069630000000000000000000000000000000000000000000000602082015250565b7f494420697320656d707479000000000000000000000000000000000000000000600082015250565b6124a5816121e4565b81146124b057600080fd5b50565b6124bc81612216565b81146124c757600080fd5b5056fea264697066735822122051b8e854abd84504aa1dae23e880a21021c2d6fec312337ea1f79923a4a9948d64736f6c63430008020033"

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

// GetEvents is a free data retrieval call binding the contract method 0x63ab919d.
//
// Solidity: function getEvents(string topicId, uint256 fromBlock) view returns((string,address[],address,uint256,uint256,bytes,bytes)[])
func (_Contract *ContractCaller) GetEvents(opts *bind.CallOpts, topicId string, fromBlock *big.Int) ([]EventLogEvent, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEvents", topicId, fromBlock)

	if err != nil {
		return *new([]EventLogEvent), err
	}

	out0 := *abi.ConvertType(out[0], new([]EventLogEvent)).(*[]EventLogEvent)

	return out0, err

}

// GetEvents is a free data retrieval call binding the contract method 0x63ab919d.
//
// Solidity: function getEvents(string topicId, uint256 fromBlock) view returns((string,address[],address,uint256,uint256,bytes,bytes)[])
func (_Contract *ContractSession) GetEvents(topicId string, fromBlock *big.Int) ([]EventLogEvent, error) {
	return _Contract.Contract.GetEvents(&_Contract.CallOpts, topicId, fromBlock)
}

// GetEvents is a free data retrieval call binding the contract method 0x63ab919d.
//
// Solidity: function getEvents(string topicId, uint256 fromBlock) view returns((string,address[],address,uint256,uint256,bytes,bytes)[])
func (_Contract *ContractCallerSession) GetEvents(topicId string, fromBlock *big.Int) ([]EventLogEvent, error) {
	return _Contract.Contract.GetEvents(&_Contract.CallOpts, topicId, fromBlock)
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

// PublishEvent is a paid mutator transaction binding the contract method 0xcc355708.
//
// Solidity: function publishEvent(string topicId, bytes id, string cid, address[] newAccounts) returns()
func (_Contract *ContractTransactor) PublishEvent(opts *bind.TransactOpts, topicId string, id []byte, cid string, newAccounts []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "publishEvent", topicId, id, cid, newAccounts)
}

// PublishEvent is a paid mutator transaction binding the contract method 0xcc355708.
//
// Solidity: function publishEvent(string topicId, bytes id, string cid, address[] newAccounts) returns()
func (_Contract *ContractSession) PublishEvent(topicId string, id []byte, cid string, newAccounts []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.PublishEvent(&_Contract.TransactOpts, topicId, id, cid, newAccounts)
}

// PublishEvent is a paid mutator transaction binding the contract method 0xcc355708.
//
// Solidity: function publishEvent(string topicId, bytes id, string cid, address[] newAccounts) returns()
func (_Contract *ContractTransactorSession) PublishEvent(topicId string, id []byte, cid string, newAccounts []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.PublishEvent(&_Contract.TransactOpts, topicId, id, cid, newAccounts)
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
	Id          []byte
	Cid         string
	NewAccounts []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEventPublished is a free log retrieval operation binding the contract event 0xbc285462bdca2ad65ad23ab3343b809e92ca2051ce3b33440b13833147aab8e7.
//
// Solidity: event EventPublished(string topicId, bytes id, string cid, address[] newAccounts)
func (_Contract *ContractFilterer) FilterEventPublished(opts *bind.FilterOpts) (*ContractEventPublishedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EventPublished")
	if err != nil {
		return nil, err
	}
	return &ContractEventPublishedIterator{contract: _Contract.contract, event: "EventPublished", logs: logs, sub: sub}, nil
}

// WatchEventPublished is a free log subscription operation binding the contract event 0xbc285462bdca2ad65ad23ab3343b809e92ca2051ce3b33440b13833147aab8e7.
//
// Solidity: event EventPublished(string topicId, bytes id, string cid, address[] newAccounts)
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

// ParseEventPublished is a log parse operation binding the contract event 0xbc285462bdca2ad65ad23ab3343b809e92ca2051ce3b33440b13833147aab8e7.
//
// Solidity: event EventPublished(string topicId, bytes id, string cid, address[] newAccounts)
func (_Contract *ContractFilterer) ParseEventPublished(log types.Log) (*ContractEventPublished, error) {
	event := new(ContractEventPublished)
	if err := _Contract.contract.UnpackLog(event, "EventPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
