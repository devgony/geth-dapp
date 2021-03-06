// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quiz

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// QuizMetaData contains all meta data concerning the Quiz contract.
var QuizMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_qn\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_ans\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkBoard\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"question\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ans\",\"type\":\"bytes32\"}],\"name\":\"sendAnswer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516104a43803806104a48339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b8382019150602082018581111561006957600080fd5b825186600182028301116401000000008211171561008657600080fd5b8083526020830192505050908051906020019080838360005b838110156100ba57808201518184015260208101905061009f565b50505050905090810190601f1680156100e75780820380516001836020036101000a031916815260200191505b50604052602001805190602001909291905050508160009080519060200190610111929190610120565b508060018190555050506101c5565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061016157805160ff191683800117855561018f565b8280016001018555821561018f579182015b8281111561018e578251825591602001919060010190610173565b5b50905061019c91906101a0565b5090565b6101c291905b808211156101be5760008160009055506001016101a6565b5090565b90565b6102d0806101d46000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806317d1653c146100465780633fad9ae01461008c57806377f46bff1461010f575b600080fd5b6100726004803603602081101561005c57600080fd5b8101908080359060200190929190505050610131565b604051808215151515815260200191505060405180910390f35b610094610147565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100d45780820151818401526020810190506100b9565b50505050905090810190601f1680156101015780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101176101e5565b604051808215151515815260200191505060405180910390f35b60006101408260015414610239565b9050919050565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101dd5780601f106101b2576101008083540402835291602001916101dd565b820191906000526020600020905b8154815290600101906020018083116101c057829003601f168201915b505050505081565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905090565b600081600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001905091905056fea265627a7a723158201bc88ce2d4032ac5b6dfdefd1c7b2eec9baacfe0351622977b248eef9dcf1fc864736f6c63430005110032",
}

// QuizABI is the input ABI used to generate the binding from.
// Deprecated: Use QuizMetaData.ABI instead.
var QuizABI = QuizMetaData.ABI

// QuizBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QuizMetaData.Bin instead.
var QuizBin = QuizMetaData.Bin

// DeployQuiz deploys a new Ethereum contract, binding an instance of Quiz to it.
func DeployQuiz(auth *bind.TransactOpts, backend bind.ContractBackend, _qn string, _ans [32]byte) (common.Address, *types.Transaction, *Quiz, error) {
	parsed, err := QuizMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QuizBin), backend, _qn, _ans)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Quiz{QuizCaller: QuizCaller{contract: contract}, QuizTransactor: QuizTransactor{contract: contract}, QuizFilterer: QuizFilterer{contract: contract}}, nil
}

// Quiz is an auto generated Go binding around an Ethereum contract.
type Quiz struct {
	QuizCaller     // Read-only binding to the contract
	QuizTransactor // Write-only binding to the contract
	QuizFilterer   // Log filterer for contract events
}

// QuizCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuizCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuizTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuizFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuizSession struct {
	Contract     *Quiz             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuizCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuizCallerSession struct {
	Contract *QuizCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QuizTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuizTransactorSession struct {
	Contract     *QuizTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuizRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuizRaw struct {
	Contract *Quiz // Generic contract binding to access the raw methods on
}

// QuizCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuizCallerRaw struct {
	Contract *QuizCaller // Generic read-only contract binding to access the raw methods on
}

// QuizTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuizTransactorRaw struct {
	Contract *QuizTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuiz creates a new instance of Quiz, bound to a specific deployed contract.
func NewQuiz(address common.Address, backend bind.ContractBackend) (*Quiz, error) {
	contract, err := bindQuiz(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quiz{QuizCaller: QuizCaller{contract: contract}, QuizTransactor: QuizTransactor{contract: contract}, QuizFilterer: QuizFilterer{contract: contract}}, nil
}

// NewQuizCaller creates a new read-only instance of Quiz, bound to a specific deployed contract.
func NewQuizCaller(address common.Address, caller bind.ContractCaller) (*QuizCaller, error) {
	contract, err := bindQuiz(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuizCaller{contract: contract}, nil
}

// NewQuizTransactor creates a new write-only instance of Quiz, bound to a specific deployed contract.
func NewQuizTransactor(address common.Address, transactor bind.ContractTransactor) (*QuizTransactor, error) {
	contract, err := bindQuiz(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuizTransactor{contract: contract}, nil
}

// NewQuizFilterer creates a new log filterer instance of Quiz, bound to a specific deployed contract.
func NewQuizFilterer(address common.Address, filterer bind.ContractFilterer) (*QuizFilterer, error) {
	contract, err := bindQuiz(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuizFilterer{contract: contract}, nil
}

// bindQuiz binds a generic wrapper to an already deployed contract.
func bindQuiz(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuizABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quiz *QuizRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quiz.Contract.QuizCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quiz *QuizRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quiz.Contract.QuizTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quiz *QuizRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quiz.Contract.QuizTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quiz *QuizCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quiz.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quiz *QuizTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quiz.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quiz *QuizTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quiz.Contract.contract.Transact(opts, method, params...)
}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() view returns(bool)
func (_Quiz *QuizCaller) CheckBoard(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Quiz.contract.Call(opts, &out, "checkBoard")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() view returns(bool)
func (_Quiz *QuizSession) CheckBoard() (bool, error) {
	return _Quiz.Contract.CheckBoard(&_Quiz.CallOpts)
}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() view returns(bool)
func (_Quiz *QuizCallerSession) CheckBoard() (bool, error) {
	return _Quiz.Contract.CheckBoard(&_Quiz.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Quiz *QuizCaller) Question(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Quiz.contract.Call(opts, &out, "question")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Quiz *QuizSession) Question() (string, error) {
	return _Quiz.Contract.Question(&_Quiz.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Quiz *QuizCallerSession) Question() (string, error) {
	return _Quiz.Contract.Question(&_Quiz.CallOpts)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizTransactor) SendAnswer(opts *bind.TransactOpts, _ans [32]byte) (*types.Transaction, error) {
	return _Quiz.contract.Transact(opts, "sendAnswer", _ans)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizSession) SendAnswer(_ans [32]byte) (*types.Transaction, error) {
	return _Quiz.Contract.SendAnswer(&_Quiz.TransactOpts, _ans)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizTransactorSession) SendAnswer(_ans [32]byte) (*types.Transaction, error) {
	return _Quiz.Contract.SendAnswer(&_Quiz.TransactOpts, _ans)
}
