// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bundler

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

// SwapCall is an auto generated low-level Go binding around an user-defined struct.
type SwapCall struct {
	Pool     common.Address
	SwapType uint8
	TokenIn  common.Address
	TokenOut common.Address
	I        *big.Int
	J        *big.Int
}

// BundlerMetaData contains all meta data concerning the Bundler contract.
var BundlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"UNISWAP_QUOTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"}],\"internalType\":\"structSwapCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getUniswapV2AmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getUniswapV3AmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BundlerABI is the input ABI used to generate the binding from.
// Deprecated: Use BundlerMetaData.ABI instead.
var BundlerABI = BundlerMetaData.ABI

// Bundler is an auto generated Go binding around an Ethereum contract.
type Bundler struct {
	BundlerCaller     // Read-only binding to the contract
	BundlerTransactor // Write-only binding to the contract
	BundlerFilterer   // Log filterer for contract events
}

// BundlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BundlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BundlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BundlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BundlerSession struct {
	Contract     *Bundler          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BundlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BundlerCallerSession struct {
	Contract *BundlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BundlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BundlerTransactorSession struct {
	Contract     *BundlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BundlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BundlerRaw struct {
	Contract *Bundler // Generic contract binding to access the raw methods on
}

// BundlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BundlerCallerRaw struct {
	Contract *BundlerCaller // Generic read-only contract binding to access the raw methods on
}

// BundlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BundlerTransactorRaw struct {
	Contract *BundlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBundler creates a new instance of Bundler, bound to a specific deployed contract.
func NewBundler(address common.Address, backend bind.ContractBackend) (*Bundler, error) {
	contract, err := bindBundler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bundler{BundlerCaller: BundlerCaller{contract: contract}, BundlerTransactor: BundlerTransactor{contract: contract}, BundlerFilterer: BundlerFilterer{contract: contract}}, nil
}

// NewBundlerCaller creates a new read-only instance of Bundler, bound to a specific deployed contract.
func NewBundlerCaller(address common.Address, caller bind.ContractCaller) (*BundlerCaller, error) {
	contract, err := bindBundler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BundlerCaller{contract: contract}, nil
}

// NewBundlerTransactor creates a new write-only instance of Bundler, bound to a specific deployed contract.
func NewBundlerTransactor(address common.Address, transactor bind.ContractTransactor) (*BundlerTransactor, error) {
	contract, err := bindBundler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BundlerTransactor{contract: contract}, nil
}

// NewBundlerFilterer creates a new log filterer instance of Bundler, bound to a specific deployed contract.
func NewBundlerFilterer(address common.Address, filterer bind.ContractFilterer) (*BundlerFilterer, error) {
	contract, err := bindBundler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BundlerFilterer{contract: contract}, nil
}

// bindBundler binds a generic wrapper to an already deployed contract.
func bindBundler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BundlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bundler *BundlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bundler.Contract.BundlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bundler *BundlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bundler.Contract.BundlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bundler *BundlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bundler.Contract.BundlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bundler *BundlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bundler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bundler *BundlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bundler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bundler *BundlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bundler.Contract.contract.Transact(opts, method, params...)
}

// UNISWAPQUOTER is a free data retrieval call binding the contract method 0xb268630b.
//
// Solidity: function UNISWAP_QUOTER() view returns(address)
func (_Bundler *BundlerCaller) UNISWAPQUOTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "UNISWAP_QUOTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPQUOTER is a free data retrieval call binding the contract method 0xb268630b.
//
// Solidity: function UNISWAP_QUOTER() view returns(address)
func (_Bundler *BundlerSession) UNISWAPQUOTER() (common.Address, error) {
	return _Bundler.Contract.UNISWAPQUOTER(&_Bundler.CallOpts)
}

// UNISWAPQUOTER is a free data retrieval call binding the contract method 0xb268630b.
//
// Solidity: function UNISWAP_QUOTER() view returns(address)
func (_Bundler *BundlerCallerSession) UNISWAPQUOTER() (common.Address, error) {
	return _Bundler.Contract.UNISWAPQUOTER(&_Bundler.CallOpts)
}

// GetUniswapV2AmountOut is a free data retrieval call binding the contract method 0x4db690d0.
//
// Solidity: function getUniswapV2AmountOut(address pool, address tokenIn, uint256 amountIn) view returns(uint256 amountOut)
func (_Bundler *BundlerCaller) GetUniswapV2AmountOut(opts *bind.CallOpts, pool common.Address, tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bundler.contract.Call(opts, &out, "getUniswapV2AmountOut", pool, tokenIn, amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUniswapV2AmountOut is a free data retrieval call binding the contract method 0x4db690d0.
//
// Solidity: function getUniswapV2AmountOut(address pool, address tokenIn, uint256 amountIn) view returns(uint256 amountOut)
func (_Bundler *BundlerSession) GetUniswapV2AmountOut(pool common.Address, tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	return _Bundler.Contract.GetUniswapV2AmountOut(&_Bundler.CallOpts, pool, tokenIn, amountIn)
}

// GetUniswapV2AmountOut is a free data retrieval call binding the contract method 0x4db690d0.
//
// Solidity: function getUniswapV2AmountOut(address pool, address tokenIn, uint256 amountIn) view returns(uint256 amountOut)
func (_Bundler *BundlerCallerSession) GetUniswapV2AmountOut(pool common.Address, tokenIn common.Address, amountIn *big.Int) (*big.Int, error) {
	return _Bundler.Contract.GetUniswapV2AmountOut(&_Bundler.CallOpts, pool, tokenIn, amountIn)
}

// GetAmountsOut is a paid mutator transaction binding the contract method 0xa9eed001.
//
// Solidity: function getAmountsOut((address,uint8,address,address,uint256,uint256)[] calls, uint256 amountIn) returns(uint256[] amountsOut)
func (_Bundler *BundlerTransactor) GetAmountsOut(opts *bind.TransactOpts, calls []SwapCall, amountIn *big.Int) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "getAmountsOut", calls, amountIn)
}

// GetAmountsOut is a paid mutator transaction binding the contract method 0xa9eed001.
//
// Solidity: function getAmountsOut((address,uint8,address,address,uint256,uint256)[] calls, uint256 amountIn) returns(uint256[] amountsOut)
func (_Bundler *BundlerSession) GetAmountsOut(calls []SwapCall, amountIn *big.Int) (*types.Transaction, error) {
	return _Bundler.Contract.GetAmountsOut(&_Bundler.TransactOpts, calls, amountIn)
}

// GetAmountsOut is a paid mutator transaction binding the contract method 0xa9eed001.
//
// Solidity: function getAmountsOut((address,uint8,address,address,uint256,uint256)[] calls, uint256 amountIn) returns(uint256[] amountsOut)
func (_Bundler *BundlerTransactorSession) GetAmountsOut(calls []SwapCall, amountIn *big.Int) (*types.Transaction, error) {
	return _Bundler.Contract.GetAmountsOut(&_Bundler.TransactOpts, calls, amountIn)
}

// GetUniswapV3AmountOut is a paid mutator transaction binding the contract method 0xb6d2cd64.
//
// Solidity: function getUniswapV3AmountOut(address pool, address tokenIn, address tokenOut, uint256 amountIn) returns(uint256 amountOut)
func (_Bundler *BundlerTransactor) GetUniswapV3AmountOut(opts *bind.TransactOpts, pool common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _Bundler.contract.Transact(opts, "getUniswapV3AmountOut", pool, tokenIn, tokenOut, amountIn)
}

// GetUniswapV3AmountOut is a paid mutator transaction binding the contract method 0xb6d2cd64.
//
// Solidity: function getUniswapV3AmountOut(address pool, address tokenIn, address tokenOut, uint256 amountIn) returns(uint256 amountOut)
func (_Bundler *BundlerSession) GetUniswapV3AmountOut(pool common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _Bundler.Contract.GetUniswapV3AmountOut(&_Bundler.TransactOpts, pool, tokenIn, tokenOut, amountIn)
}

// GetUniswapV3AmountOut is a paid mutator transaction binding the contract method 0xb6d2cd64.
//
// Solidity: function getUniswapV3AmountOut(address pool, address tokenIn, address tokenOut, uint256 amountIn) returns(uint256 amountOut)
func (_Bundler *BundlerTransactorSession) GetUniswapV3AmountOut(pool common.Address, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int) (*types.Transaction, error) {
	return _Bundler.Contract.GetUniswapV3AmountOut(&_Bundler.TransactOpts, pool, tokenIn, tokenOut, amountIn)
}
