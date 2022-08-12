// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cryptoswap

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

// CurveCryptoSwapMetaData contains all meta data concerning the CurveCryptoSwap contract.
var CurveCryptoSwapMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[3]\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[3]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_index\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewParameters\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampAgamma\",\"inputs\":[{\"name\":\"initial_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"current_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"current_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ClaimAdminFee\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true},{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"admin_fee_receiver\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"gamma\",\"type\":\"uint256\"},{\"name\":\"mid_fee\",\"type\":\"uint256\"},{\"name\":\"out_fee\",\"type\":\"uint256\"},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"fee_gamma\",\"type\":\"uint256\"},{\"name\":\"adjustment_step\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"ma_half_time\",\"type\":\"uint256\"},{\"name\":\"initial_prices\",\"type\":\"uint256[2]\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3361},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_scale\",\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3391},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices\",\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3421},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":468},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":582},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":597},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":11991},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":21673},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_calc\",\"inputs\":[{\"name\":\"xp\",\"type\":\"uint256[3]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":11096},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":11582},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3122},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_fee\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"xp\",\"type\":\"uint256[3]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":26582},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":738687},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[3]\"}],\"outputs\":[],\"gas\":233981},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3429},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":13432},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":648579},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_admin_fees\",\"inputs\":[],\"outputs\":[],\"gas\":389808},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A_gamma\",\"inputs\":[{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"future_gamma\",\"type\":\"uint256\"},{\"name\":\"future_time\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":163102},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A_gamma\",\"inputs\":[],\"outputs\":[],\"gas\":157247},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_parameters\",\"inputs\":[{\"name\":\"_new_mid_fee\",\"type\":\"uint256\"},{\"name\":\"_new_out_fee\",\"type\":\"uint256\"},{\"name\":\"_new_admin_fee\",\"type\":\"uint256\"},{\"name\":\"_new_fee_gamma\",\"type\":\"uint256\"},{\"name\":\"_new_allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"_new_adjustment_step\",\"type\":\"uint256\"},{\"name\":\"_new_ma_half_time\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":306190},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_parameters\",\"inputs\":[],\"outputs\":[],\"gas\":683438},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revert_new_parameters\",\"inputs\":[],\"outputs\":[],\"gas\":23222},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"outputs\":[],\"gas\":77260},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_transfer_ownership\",\"inputs\":[],\"outputs\":[],\"gas\":65937},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revert_transfer_ownership\",\"inputs\":[],\"outputs\":[],\"gas\":23312},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"kill_me\",\"inputs\":[],\"outputs\":[],\"gas\":40535},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"unkill_me\",\"inputs\":[],\"outputs\":[],\"gas\":23372},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_admin_fee_receiver\",\"inputs\":[{\"name\":\"_admin_fee_receiver\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38505},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices_timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3378},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3408},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3438},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3468},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3498},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3528},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3558},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3588},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3618},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3648},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3678},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3708},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_ma_half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3738},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3768},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3798},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3828},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3858},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3888},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3918},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4057},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3978},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":4008},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":4038},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4068},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit_a\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4098},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4128},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_killed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":4158},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"kill_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4188},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"transfer_ownership_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4218},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_actions_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4248},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee_receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":4278}]",
}

// CurveCryptoSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveCryptoSwapMetaData.ABI instead.
var CurveCryptoSwapABI = CurveCryptoSwapMetaData.ABI

// CurveCryptoSwap is an auto generated Go binding around an Ethereum contract.
type CurveCryptoSwap struct {
	CurveCryptoSwapCaller     // Read-only binding to the contract
	CurveCryptoSwapTransactor // Write-only binding to the contract
	CurveCryptoSwapFilterer   // Log filterer for contract events
}

// CurveCryptoSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveCryptoSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveCryptoSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveCryptoSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveCryptoSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveCryptoSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveCryptoSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveCryptoSwapSession struct {
	Contract     *CurveCryptoSwap  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveCryptoSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveCryptoSwapCallerSession struct {
	Contract *CurveCryptoSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CurveCryptoSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveCryptoSwapTransactorSession struct {
	Contract     *CurveCryptoSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CurveCryptoSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveCryptoSwapRaw struct {
	Contract *CurveCryptoSwap // Generic contract binding to access the raw methods on
}

// CurveCryptoSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveCryptoSwapCallerRaw struct {
	Contract *CurveCryptoSwapCaller // Generic read-only contract binding to access the raw methods on
}

// CurveCryptoSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveCryptoSwapTransactorRaw struct {
	Contract *CurveCryptoSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveCryptoSwap creates a new instance of CurveCryptoSwap, bound to a specific deployed contract.
func NewCurveCryptoSwap(address common.Address, backend bind.ContractBackend) (*CurveCryptoSwap, error) {
	contract, err := bindCurveCryptoSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwap{CurveCryptoSwapCaller: CurveCryptoSwapCaller{contract: contract}, CurveCryptoSwapTransactor: CurveCryptoSwapTransactor{contract: contract}, CurveCryptoSwapFilterer: CurveCryptoSwapFilterer{contract: contract}}, nil
}

// NewCurveCryptoSwapCaller creates a new read-only instance of CurveCryptoSwap, bound to a specific deployed contract.
func NewCurveCryptoSwapCaller(address common.Address, caller bind.ContractCaller) (*CurveCryptoSwapCaller, error) {
	contract, err := bindCurveCryptoSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapCaller{contract: contract}, nil
}

// NewCurveCryptoSwapTransactor creates a new write-only instance of CurveCryptoSwap, bound to a specific deployed contract.
func NewCurveCryptoSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveCryptoSwapTransactor, error) {
	contract, err := bindCurveCryptoSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapTransactor{contract: contract}, nil
}

// NewCurveCryptoSwapFilterer creates a new log filterer instance of CurveCryptoSwap, bound to a specific deployed contract.
func NewCurveCryptoSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveCryptoSwapFilterer, error) {
	contract, err := bindCurveCryptoSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFilterer{contract: contract}, nil
}

// bindCurveCryptoSwap binds a generic wrapper to an already deployed contract.
func bindCurveCryptoSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CurveCryptoSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveCryptoSwap *CurveCryptoSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveCryptoSwap.Contract.CurveCryptoSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveCryptoSwap *CurveCryptoSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.CurveCryptoSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveCryptoSwap *CurveCryptoSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.CurveCryptoSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveCryptoSwap *CurveCryptoSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveCryptoSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveCryptoSwap *CurveCryptoSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveCryptoSwap *CurveCryptoSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) A() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.A(&_CurveCryptoSwap.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) A() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.A(&_CurveCryptoSwap.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) D(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "D")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) D() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.D(&_CurveCryptoSwap.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) D() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.D(&_CurveCryptoSwap.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) AdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) AdjustmentStep() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.AdjustmentStep(&_CurveCryptoSwap.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) AdjustmentStep() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.AdjustmentStep(&_CurveCryptoSwap.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.AdminActionsDeadline(&_CurveCryptoSwap.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.AdminActionsDeadline(&_CurveCryptoSwap.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) AdminFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.AdminFee(&_CurveCryptoSwap.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) AdminFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.AdminFee(&_CurveCryptoSwap.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) AdminFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "admin_fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapSession) AdminFeeReceiver() (common.Address, error) {
	return _CurveCryptoSwap.Contract.AdminFeeReceiver(&_CurveCryptoSwap.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) AdminFeeReceiver() (common.Address, error) {
	return _CurveCryptoSwap.Contract.AdminFeeReceiver(&_CurveCryptoSwap.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) AllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) AllowedExtraProfit() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.AllowedExtraProfit(&_CurveCryptoSwap.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) AllowedExtraProfit() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.AllowedExtraProfit(&_CurveCryptoSwap.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.Balances(&_CurveCryptoSwap.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.Balances(&_CurveCryptoSwap.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) CalcTokenAmount(opts *bind.CallOpts, amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "calc_token_amount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.CalcTokenAmount(&_CurveCryptoSwap.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.CalcTokenAmount(&_CurveCryptoSwap.CallOpts, amounts, deposit)
}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) CalcTokenFee(opts *bind.CallOpts, amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "calc_token_fee", amounts, xp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) CalcTokenFee(amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.CalcTokenFee(&_CurveCryptoSwap.CallOpts, amounts, xp)
}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) CalcTokenFee(amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.CalcTokenFee(&_CurveCryptoSwap.CallOpts, amounts, xp)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) CalcWithdrawOneCoin(opts *bind.CallOpts, token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "calc_withdraw_one_coin", token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.CalcWithdrawOneCoin(&_CurveCryptoSwap.CallOpts, token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.CalcWithdrawOneCoin(&_CurveCryptoSwap.CallOpts, token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) Coins(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "coins", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapSession) Coins(i *big.Int) (common.Address, error) {
	return _CurveCryptoSwap.Contract.Coins(&_CurveCryptoSwap.CallOpts, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) Coins(i *big.Int) (common.Address, error) {
	return _CurveCryptoSwap.Contract.Coins(&_CurveCryptoSwap.CallOpts, i)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) Fee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.Fee(&_CurveCryptoSwap.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) Fee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.Fee(&_CurveCryptoSwap.CallOpts)
}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FeeCalc(opts *bind.CallOpts, xp [3]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "fee_calc", xp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FeeCalc(xp [3]*big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FeeCalc(&_CurveCryptoSwap.CallOpts, xp)
}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FeeCalc(xp [3]*big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FeeCalc(&_CurveCryptoSwap.CallOpts, xp)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FeeGamma() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FeeGamma(&_CurveCryptoSwap.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FeeGamma() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FeeGamma(&_CurveCryptoSwap.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FutureAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "future_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FutureAGamma() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureAGamma(&_CurveCryptoSwap.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FutureAGamma() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureAGamma(&_CurveCryptoSwap.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FutureAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "future_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FutureAGammaTime() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureAGammaTime(&_CurveCryptoSwap.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FutureAGammaTime() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureAGammaTime(&_CurveCryptoSwap.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FutureAdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "future_adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FutureAdjustmentStep() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureAdjustmentStep(&_CurveCryptoSwap.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FutureAdjustmentStep() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureAdjustmentStep(&_CurveCryptoSwap.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FutureAdminFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureAdminFee(&_CurveCryptoSwap.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FutureAdminFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureAdminFee(&_CurveCryptoSwap.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FutureAllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "future_allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureAllowedExtraProfit(&_CurveCryptoSwap.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureAllowedExtraProfit(&_CurveCryptoSwap.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FutureFeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "future_fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FutureFeeGamma() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureFeeGamma(&_CurveCryptoSwap.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FutureFeeGamma() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureFeeGamma(&_CurveCryptoSwap.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FutureMaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "future_ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FutureMaHalfTime() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureMaHalfTime(&_CurveCryptoSwap.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FutureMaHalfTime() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureMaHalfTime(&_CurveCryptoSwap.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FutureMidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "future_mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FutureMidFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureMidFee(&_CurveCryptoSwap.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FutureMidFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureMidFee(&_CurveCryptoSwap.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FutureOutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "future_out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FutureOutFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureOutFee(&_CurveCryptoSwap.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FutureOutFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.FutureOutFee(&_CurveCryptoSwap.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapSession) FutureOwner() (common.Address, error) {
	return _CurveCryptoSwap.Contract.FutureOwner(&_CurveCryptoSwap.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) FutureOwner() (common.Address, error) {
	return _CurveCryptoSwap.Contract.FutureOwner(&_CurveCryptoSwap.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) Gamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) Gamma() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.Gamma(&_CurveCryptoSwap.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) Gamma() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.Gamma(&_CurveCryptoSwap.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.GetDy(&_CurveCryptoSwap.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.GetDy(&_CurveCryptoSwap.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.GetVirtualPrice(&_CurveCryptoSwap.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.GetVirtualPrice(&_CurveCryptoSwap.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) InitialAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "initial_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) InitialAGamma() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.InitialAGamma(&_CurveCryptoSwap.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) InitialAGamma() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.InitialAGamma(&_CurveCryptoSwap.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) InitialAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "initial_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) InitialAGammaTime() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.InitialAGammaTime(&_CurveCryptoSwap.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) InitialAGammaTime() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.InitialAGammaTime(&_CurveCryptoSwap.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "is_killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveCryptoSwap *CurveCryptoSwapSession) IsKilled() (bool, error) {
	return _CurveCryptoSwap.Contract.IsKilled(&_CurveCryptoSwap.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) IsKilled() (bool, error) {
	return _CurveCryptoSwap.Contract.IsKilled(&_CurveCryptoSwap.CallOpts)
}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) KillDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "kill_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) KillDeadline() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.KillDeadline(&_CurveCryptoSwap.CallOpts)
}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) KillDeadline() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.KillDeadline(&_CurveCryptoSwap.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) LastPrices(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "last_prices", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) LastPrices(k *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.LastPrices(&_CurveCryptoSwap.CallOpts, k)
}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) LastPrices(k *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.LastPrices(&_CurveCryptoSwap.CallOpts, k)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) LastPricesTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "last_prices_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) LastPricesTimestamp() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.LastPricesTimestamp(&_CurveCryptoSwap.CallOpts)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) LastPricesTimestamp() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.LastPricesTimestamp(&_CurveCryptoSwap.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) MaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) MaHalfTime() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.MaHalfTime(&_CurveCryptoSwap.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) MaHalfTime() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.MaHalfTime(&_CurveCryptoSwap.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) MidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) MidFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.MidFee(&_CurveCryptoSwap.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) MidFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.MidFee(&_CurveCryptoSwap.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) OutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) OutFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.OutFee(&_CurveCryptoSwap.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) OutFee() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.OutFee(&_CurveCryptoSwap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapSession) Owner() (common.Address, error) {
	return _CurveCryptoSwap.Contract.Owner(&_CurveCryptoSwap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) Owner() (common.Address, error) {
	return _CurveCryptoSwap.Contract.Owner(&_CurveCryptoSwap.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) PriceOracle(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "price_oracle", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) PriceOracle(k *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.PriceOracle(&_CurveCryptoSwap.CallOpts, k)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) PriceOracle(k *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.PriceOracle(&_CurveCryptoSwap.CallOpts, k)
}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) PriceScale(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "price_scale", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) PriceScale(k *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.PriceScale(&_CurveCryptoSwap.CallOpts, k)
}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) PriceScale(k *big.Int) (*big.Int, error) {
	return _CurveCryptoSwap.Contract.PriceScale(&_CurveCryptoSwap.CallOpts, k)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapSession) Token() (common.Address, error) {
	return _CurveCryptoSwap.Contract.Token(&_CurveCryptoSwap.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) Token() (common.Address, error) {
	return _CurveCryptoSwap.Contract.Token(&_CurveCryptoSwap.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.TransferOwnershipDeadline(&_CurveCryptoSwap.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.TransferOwnershipDeadline(&_CurveCryptoSwap.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) VirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) VirtualPrice() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.VirtualPrice(&_CurveCryptoSwap.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) VirtualPrice() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.VirtualPrice(&_CurveCryptoSwap.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) XcpProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "xcp_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) XcpProfit() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.XcpProfit(&_CurveCryptoSwap.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) XcpProfit() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.XcpProfit(&_CurveCryptoSwap.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCaller) XcpProfitA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwap.contract.Call(opts, &out, "xcp_profit_a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapSession) XcpProfitA() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.XcpProfitA(&_CurveCryptoSwap.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_CurveCryptoSwap *CurveCryptoSwapCallerSession) XcpProfitA() (*big.Int, error) {
	return _CurveCryptoSwap.Contract.XcpProfitA(&_CurveCryptoSwap.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) AddLiquidity(opts *bind.TransactOpts, amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.AddLiquidity(&_CurveCryptoSwap.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.AddLiquidity(&_CurveCryptoSwap.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) ApplyNewParameters() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.ApplyNewParameters(&_CurveCryptoSwap.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.ApplyNewParameters(&_CurveCryptoSwap.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.ApplyTransferOwnership(&_CurveCryptoSwap.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.ApplyTransferOwnership(&_CurveCryptoSwap.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) ClaimAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "claim_admin_fees")
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) ClaimAdminFees() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.ClaimAdminFees(&_CurveCryptoSwap.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) ClaimAdminFees() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.ClaimAdminFees(&_CurveCryptoSwap.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) CommitNewParameters(opts *bind.TransactOpts, _new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "commit_new_parameters", _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.CommitNewParameters(&_CurveCryptoSwap.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.CommitNewParameters(&_CurveCryptoSwap.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.CommitTransferOwnership(&_CurveCryptoSwap.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.CommitTransferOwnership(&_CurveCryptoSwap.TransactOpts, _owner)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.Exchange(&_CurveCryptoSwap.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.Exchange(&_CurveCryptoSwap.TransactOpts, i, j, dx, min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "exchange0", i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.Exchange0(&_CurveCryptoSwap.TransactOpts, i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.Exchange0(&_CurveCryptoSwap.TransactOpts, i, j, dx, min_dy, use_eth)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) KillMe() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.KillMe(&_CurveCryptoSwap.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) KillMe() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.KillMe(&_CurveCryptoSwap.TransactOpts)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) RampAGamma(opts *bind.TransactOpts, future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "ramp_A_gamma", future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.RampAGamma(&_CurveCryptoSwap.TransactOpts, future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.RampAGamma(&_CurveCryptoSwap.TransactOpts, future_A, future_gamma, future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.RemoveLiquidity(&_CurveCryptoSwap.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.RemoveLiquidity(&_CurveCryptoSwap.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "remove_liquidity_one_coin", token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.RemoveLiquidityOneCoin(&_CurveCryptoSwap.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.RemoveLiquidityOneCoin(&_CurveCryptoSwap.TransactOpts, token_amount, i, min_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.RevertNewParameters(&_CurveCryptoSwap.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.RevertNewParameters(&_CurveCryptoSwap.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.RevertTransferOwnership(&_CurveCryptoSwap.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.RevertTransferOwnership(&_CurveCryptoSwap.TransactOpts)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) SetAdminFeeReceiver(opts *bind.TransactOpts, _admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "set_admin_fee_receiver", _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.SetAdminFeeReceiver(&_CurveCryptoSwap.TransactOpts, _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.SetAdminFeeReceiver(&_CurveCryptoSwap.TransactOpts, _admin_fee_receiver)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) StopRampAGamma(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "stop_ramp_A_gamma")
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) StopRampAGamma() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.StopRampAGamma(&_CurveCryptoSwap.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) StopRampAGamma() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.StopRampAGamma(&_CurveCryptoSwap.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) UnkillMe() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.UnkillMe(&_CurveCryptoSwap.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) UnkillMe() (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.UnkillMe(&_CurveCryptoSwap.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CurveCryptoSwap.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CurveCryptoSwap *CurveCryptoSwapSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.Fallback(&_CurveCryptoSwap.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CurveCryptoSwap *CurveCryptoSwapTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CurveCryptoSwap.Contract.Fallback(&_CurveCryptoSwap.TransactOpts, calldata)
}

// CurveCryptoSwapAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapAddLiquidityIterator struct {
	Event *CurveCryptoSwapAddLiquidity // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapAddLiquidity)
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
		it.Event = new(CurveCryptoSwapAddLiquidity)
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
func (it *CurveCryptoSwapAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapAddLiquidity represents a AddLiquidity event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapAddLiquidity struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fee          *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x96b486485420b963edd3fdec0b0195730035600feb7de6f544383d7950fa97ee.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256 fee, uint256 token_supply)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurveCryptoSwapAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapAddLiquidityIterator{contract: _CurveCryptoSwap.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x96b486485420b963edd3fdec0b0195730035600feb7de6f544383d7950fa97ee.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256 fee, uint256 token_supply)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapAddLiquidity)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x96b486485420b963edd3fdec0b0195730035600feb7de6f544383d7950fa97ee.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256 fee, uint256 token_supply)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseAddLiquidity(log types.Log) (*CurveCryptoSwapAddLiquidity, error) {
	event := new(CurveCryptoSwapAddLiquidity)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapClaimAdminFeeIterator is returned from FilterClaimAdminFee and is used to iterate over the raw logs and unpacked data for ClaimAdminFee events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapClaimAdminFeeIterator struct {
	Event *CurveCryptoSwapClaimAdminFee // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapClaimAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapClaimAdminFee)
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
		it.Event = new(CurveCryptoSwapClaimAdminFee)
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
func (it *CurveCryptoSwapClaimAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapClaimAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapClaimAdminFee represents a ClaimAdminFee event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapClaimAdminFee struct {
	Admin  common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimAdminFee is a free log retrieval operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterClaimAdminFee(opts *bind.FilterOpts, admin []common.Address) (*CurveCryptoSwapClaimAdminFeeIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapClaimAdminFeeIterator{contract: _CurveCryptoSwap.contract, event: "ClaimAdminFee", logs: logs, sub: sub}, nil
}

// WatchClaimAdminFee is a free log subscription operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchClaimAdminFee(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapClaimAdminFee, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapClaimAdminFee)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
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

// ParseClaimAdminFee is a log parse operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseClaimAdminFee(log types.Log) (*CurveCryptoSwapClaimAdminFee, error) {
	event := new(CurveCryptoSwapClaimAdminFee)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapCommitNewAdminIterator struct {
	Event *CurveCryptoSwapCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapCommitNewAdmin)
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
		it.Event = new(CurveCryptoSwapCommitNewAdmin)
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
func (it *CurveCryptoSwapCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapCommitNewAdmin represents a CommitNewAdmin event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*CurveCryptoSwapCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapCommitNewAdminIterator{contract: _CurveCryptoSwap.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapCommitNewAdmin)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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

// ParseCommitNewAdmin is a log parse operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseCommitNewAdmin(log types.Log) (*CurveCryptoSwapCommitNewAdmin, error) {
	event := new(CurveCryptoSwapCommitNewAdmin)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapCommitNewParametersIterator is returned from FilterCommitNewParameters and is used to iterate over the raw logs and unpacked data for CommitNewParameters events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapCommitNewParametersIterator struct {
	Event *CurveCryptoSwapCommitNewParameters // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapCommitNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapCommitNewParameters)
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
		it.Event = new(CurveCryptoSwapCommitNewParameters)
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
func (it *CurveCryptoSwapCommitNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapCommitNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapCommitNewParameters represents a CommitNewParameters event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapCommitNewParameters struct {
	Deadline           *big.Int
	AdminFee           *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaHalfTime         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCommitNewParameters is a free log retrieval operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterCommitNewParameters(opts *bind.FilterOpts, deadline []*big.Int) (*CurveCryptoSwapCommitNewParametersIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapCommitNewParametersIterator{contract: _CurveCryptoSwap.contract, event: "CommitNewParameters", logs: logs, sub: sub}, nil
}

// WatchCommitNewParameters is a free log subscription operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchCommitNewParameters(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapCommitNewParameters, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapCommitNewParameters)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
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

// ParseCommitNewParameters is a log parse operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseCommitNewParameters(log types.Log) (*CurveCryptoSwapCommitNewParameters, error) {
	event := new(CurveCryptoSwapCommitNewParameters)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapNewAdminIterator struct {
	Event *CurveCryptoSwapNewAdmin // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapNewAdmin)
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
		it.Event = new(CurveCryptoSwapNewAdmin)
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
func (it *CurveCryptoSwapNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapNewAdmin represents a NewAdmin event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*CurveCryptoSwapNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapNewAdminIterator{contract: _CurveCryptoSwap.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapNewAdmin)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseNewAdmin(log types.Log) (*CurveCryptoSwapNewAdmin, error) {
	event := new(CurveCryptoSwapNewAdmin)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapNewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapNewParametersIterator struct {
	Event *CurveCryptoSwapNewParameters // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapNewParameters)
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
		it.Event = new(CurveCryptoSwapNewParameters)
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
func (it *CurveCryptoSwapNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapNewParameters represents a NewParameters event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapNewParameters struct {
	AdminFee           *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaHalfTime         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterNewParameters(opts *bind.FilterOpts) (*CurveCryptoSwapNewParametersIterator, error) {

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapNewParametersIterator{contract: _CurveCryptoSwap.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapNewParameters) (event.Subscription, error) {

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapNewParameters)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "NewParameters", log); err != nil {
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

// ParseNewParameters is a log parse operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseNewParameters(log types.Log) (*CurveCryptoSwapNewParameters, error) {
	event := new(CurveCryptoSwapNewParameters)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapRampAgammaIterator is returned from FilterRampAgamma and is used to iterate over the raw logs and unpacked data for RampAgamma events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapRampAgammaIterator struct {
	Event *CurveCryptoSwapRampAgamma // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapRampAgammaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapRampAgamma)
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
		it.Event = new(CurveCryptoSwapRampAgamma)
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
func (it *CurveCryptoSwapRampAgammaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapRampAgammaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapRampAgamma represents a RampAgamma event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapRampAgamma struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialGamma *big.Int
	FutureGamma  *big.Int
	InitialTime  *big.Int
	FutureTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRampAgamma is a free log retrieval operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterRampAgamma(opts *bind.FilterOpts) (*CurveCryptoSwapRampAgammaIterator, error) {

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapRampAgammaIterator{contract: _CurveCryptoSwap.contract, event: "RampAgamma", logs: logs, sub: sub}, nil
}

// WatchRampAgamma is a free log subscription operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchRampAgamma(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapRampAgamma) (event.Subscription, error) {

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapRampAgamma)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "RampAgamma", log); err != nil {
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

// ParseRampAgamma is a log parse operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseRampAgamma(log types.Log) (*CurveCryptoSwapRampAgamma, error) {
	event := new(CurveCryptoSwapRampAgamma)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "RampAgamma", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapRemoveLiquidityIterator struct {
	Event *CurveCryptoSwapRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapRemoveLiquidity)
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
		it.Event = new(CurveCryptoSwapRemoveLiquidity)
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
func (it *CurveCryptoSwapRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapRemoveLiquidity represents a RemoveLiquidity event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xd6cc314a0b1e3b2579f8e64248e82434072e8271290eef8ad0886709304195f5.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256 token_supply)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*CurveCryptoSwapRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapRemoveLiquidityIterator{contract: _CurveCryptoSwap.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xd6cc314a0b1e3b2579f8e64248e82434072e8271290eef8ad0886709304195f5.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256 token_supply)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapRemoveLiquidity)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xd6cc314a0b1e3b2579f8e64248e82434072e8271290eef8ad0886709304195f5.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256 token_supply)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseRemoveLiquidity(log types.Log) (*CurveCryptoSwapRemoveLiquidity, error) {
	event := new(CurveCryptoSwapRemoveLiquidity)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapRemoveLiquidityOneIterator struct {
	Event *CurveCryptoSwapRemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapRemoveLiquidityOne)
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
		it.Event = new(CurveCryptoSwapRemoveLiquidityOne)
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
func (it *CurveCryptoSwapRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapRemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinIndex   *big.Int
	CoinAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*CurveCryptoSwapRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapRemoveLiquidityOneIterator{contract: _CurveCryptoSwap.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapRemoveLiquidityOne)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseRemoveLiquidityOne(log types.Log) (*CurveCryptoSwapRemoveLiquidityOne, error) {
	event := new(CurveCryptoSwapRemoveLiquidityOne)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapStopRampAIterator struct {
	Event *CurveCryptoSwapStopRampA // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapStopRampA)
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
		it.Event = new(CurveCryptoSwapStopRampA)
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
func (it *CurveCryptoSwapStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapStopRampA represents a StopRampA event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapStopRampA struct {
	CurrentA     *big.Int
	CurrentGamma *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterStopRampA(opts *bind.FilterOpts) (*CurveCryptoSwapStopRampAIterator, error) {

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapStopRampAIterator{contract: _CurveCryptoSwap.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapStopRampA) (event.Subscription, error) {

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapStopRampA)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "StopRampA", log); err != nil {
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

// ParseStopRampA is a log parse operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseStopRampA(log types.Log) (*CurveCryptoSwapStopRampA, error) {
	event := new(CurveCryptoSwapStopRampA)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapTokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the CurveCryptoSwap contract.
type CurveCryptoSwapTokenExchangeIterator struct {
	Event *CurveCryptoSwapTokenExchange // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapTokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapTokenExchange)
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
		it.Event = new(CurveCryptoSwapTokenExchange)
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
func (it *CurveCryptoSwapTokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapTokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapTokenExchange represents a TokenExchange event raised by the CurveCryptoSwap contract.
type CurveCryptoSwapTokenExchange struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*CurveCryptoSwapTokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapTokenExchangeIterator{contract: _CurveCryptoSwap.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapTokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _CurveCryptoSwap.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapTokenExchange)
				if err := _CurveCryptoSwap.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_CurveCryptoSwap *CurveCryptoSwapFilterer) ParseTokenExchange(log types.Log) (*CurveCryptoSwapTokenExchange, error) {
	event := new(CurveCryptoSwapTokenExchange)
	if err := _CurveCryptoSwap.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
