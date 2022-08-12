// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package query

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

// PairQuery is an auto generated low-level Go binding around an user-defined struct.
type PairQuery struct {
	Factory common.Address
	TokenA  common.Address
	TokenB  common.Address
}

// UniswapQueryMetaData contains all meta data concerning the UniswapQuery contract.
var UniswapQueryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"contractUniswapV2Factory\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"internalType\":\"structPairQuery\",\"name\":\"query\",\"type\":\"tuple\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractUniswapV2Factory\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"internalType\":\"structPairQuery[]\",\"name\":\"queries\",\"type\":\"tuple[]\"}],\"name\":\"getPairs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV2Pair[]\",\"name\":\"_pairs\",\"type\":\"address[]\"}],\"name\":\"getReservesByPairs\",\"outputs\":[{\"internalType\":\"uint256[2][]\",\"name\":\"\",\"type\":\"uint256[2][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapQueryABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapQueryMetaData.ABI instead.
var UniswapQueryABI = UniswapQueryMetaData.ABI

// UniswapQuery is an auto generated Go binding around an Ethereum contract.
type UniswapQuery struct {
	UniswapQueryCaller     // Read-only binding to the contract
	UniswapQueryTransactor // Write-only binding to the contract
	UniswapQueryFilterer   // Log filterer for contract events
}

// UniswapQueryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapQueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapQueryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapQueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapQueryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapQueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapQuerySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapQuerySession struct {
	Contract     *UniswapQuery     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapQueryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapQueryCallerSession struct {
	Contract *UniswapQueryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UniswapQueryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapQueryTransactorSession struct {
	Contract     *UniswapQueryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UniswapQueryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapQueryRaw struct {
	Contract *UniswapQuery // Generic contract binding to access the raw methods on
}

// UniswapQueryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapQueryCallerRaw struct {
	Contract *UniswapQueryCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapQueryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapQueryTransactorRaw struct {
	Contract *UniswapQueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapQuery creates a new instance of UniswapQuery, bound to a specific deployed contract.
func NewUniswapQuery(address common.Address, backend bind.ContractBackend) (*UniswapQuery, error) {
	contract, err := bindUniswapQuery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapQuery{UniswapQueryCaller: UniswapQueryCaller{contract: contract}, UniswapQueryTransactor: UniswapQueryTransactor{contract: contract}, UniswapQueryFilterer: UniswapQueryFilterer{contract: contract}}, nil
}

// NewUniswapQueryCaller creates a new read-only instance of UniswapQuery, bound to a specific deployed contract.
func NewUniswapQueryCaller(address common.Address, caller bind.ContractCaller) (*UniswapQueryCaller, error) {
	contract, err := bindUniswapQuery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapQueryCaller{contract: contract}, nil
}

// NewUniswapQueryTransactor creates a new write-only instance of UniswapQuery, bound to a specific deployed contract.
func NewUniswapQueryTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapQueryTransactor, error) {
	contract, err := bindUniswapQuery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapQueryTransactor{contract: contract}, nil
}

// NewUniswapQueryFilterer creates a new log filterer instance of UniswapQuery, bound to a specific deployed contract.
func NewUniswapQueryFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapQueryFilterer, error) {
	contract, err := bindUniswapQuery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapQueryFilterer{contract: contract}, nil
}

// bindUniswapQuery binds a generic wrapper to an already deployed contract.
func bindUniswapQuery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapQueryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapQuery *UniswapQueryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapQuery.Contract.UniswapQueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapQuery *UniswapQueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapQuery.Contract.UniswapQueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapQuery *UniswapQueryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapQuery.Contract.UniswapQueryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapQuery *UniswapQueryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapQuery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapQuery *UniswapQueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapQuery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapQuery *UniswapQueryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapQuery.Contract.contract.Transact(opts, method, params...)
}

// GetPair is a free data retrieval call binding the contract method 0x1947b1fd.
//
// Solidity: function getPair((address,address,address) query) view returns(address)
func (_UniswapQuery *UniswapQueryCaller) GetPair(opts *bind.CallOpts, query PairQuery) (common.Address, error) {
	var out []interface{}
	err := _UniswapQuery.contract.Call(opts, &out, "getPair", query)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0x1947b1fd.
//
// Solidity: function getPair((address,address,address) query) view returns(address)
func (_UniswapQuery *UniswapQuerySession) GetPair(query PairQuery) (common.Address, error) {
	return _UniswapQuery.Contract.GetPair(&_UniswapQuery.CallOpts, query)
}

// GetPair is a free data retrieval call binding the contract method 0x1947b1fd.
//
// Solidity: function getPair((address,address,address) query) view returns(address)
func (_UniswapQuery *UniswapQueryCallerSession) GetPair(query PairQuery) (common.Address, error) {
	return _UniswapQuery.Contract.GetPair(&_UniswapQuery.CallOpts, query)
}

// GetPairs is a free data retrieval call binding the contract method 0xf5e955c9.
//
// Solidity: function getPairs((address,address,address)[] queries) view returns(address[])
func (_UniswapQuery *UniswapQueryCaller) GetPairs(opts *bind.CallOpts, queries []PairQuery) ([]common.Address, error) {
	var out []interface{}
	err := _UniswapQuery.contract.Call(opts, &out, "getPairs", queries)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPairs is a free data retrieval call binding the contract method 0xf5e955c9.
//
// Solidity: function getPairs((address,address,address)[] queries) view returns(address[])
func (_UniswapQuery *UniswapQuerySession) GetPairs(queries []PairQuery) ([]common.Address, error) {
	return _UniswapQuery.Contract.GetPairs(&_UniswapQuery.CallOpts, queries)
}

// GetPairs is a free data retrieval call binding the contract method 0xf5e955c9.
//
// Solidity: function getPairs((address,address,address)[] queries) view returns(address[])
func (_UniswapQuery *UniswapQueryCallerSession) GetPairs(queries []PairQuery) ([]common.Address, error) {
	return _UniswapQuery.Contract.GetPairs(&_UniswapQuery.CallOpts, queries)
}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[2][])
func (_UniswapQuery *UniswapQueryCaller) GetReservesByPairs(opts *bind.CallOpts, _pairs []common.Address) ([][2]*big.Int, error) {
	var out []interface{}
	err := _UniswapQuery.contract.Call(opts, &out, "getReservesByPairs", _pairs)

	if err != nil {
		return *new([][2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][2]*big.Int)).(*[][2]*big.Int)

	return out0, err

}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[2][])
func (_UniswapQuery *UniswapQuerySession) GetReservesByPairs(_pairs []common.Address) ([][2]*big.Int, error) {
	return _UniswapQuery.Contract.GetReservesByPairs(&_UniswapQuery.CallOpts, _pairs)
}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[2][])
func (_UniswapQuery *UniswapQueryCallerSession) GetReservesByPairs(_pairs []common.Address) ([][2]*big.Int, error) {
	return _UniswapQuery.Contract.GetReservesByPairs(&_UniswapQuery.CallOpts, _pairs)
}
