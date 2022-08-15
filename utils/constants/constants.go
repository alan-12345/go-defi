package constants

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NetworkData struct {
	RPC          string
	Client       *ethclient.Client
	Headers      chan *types.Header
	Subscription ethereum.Subscription
}

type Token struct {
	Address  common.Address
	Decimals int64
	Size     *big.Int
}

type Pool struct {
	Name     string
	Tokens   []Token
	Protocol string
	SwapType uint8
}

type Compound struct {
	Unitroller common.Address
	StartBlock *big.Int
}

var (
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	Inf         = new(big.Float).SetInf(false)
	NegInf      = new(big.Float).SetInf(true)
	One         = new(big.Float).SetFloat64(1)
	NegOne      = new(big.Float).SetFloat64(-1)
	Zero        = new(big.Float).SetFloat64(0)
	TenInt      = big.NewInt(10)
	SwapTypes   = map[string]uint8{
		"UniswapV2":             0,
		"UniswapV3":             1,
		"CurveStableBase":       2,
		"CurveStableUnderlying": 3,
		"CurveCryptoBase":       4,
		"CurveCryptoUnderlying": 5,
		"CurveMetaPool":         6,
	}
	ReallyBigInt   = new(big.Int).Sub(BigPow(2, 128), big.NewInt(1))
	TelegramUrl    = "https://api.telegram.org/bot5001621564:AAGP-uUCxGgFlnwDqGuviNqDDLE91VREhao/sendMessage"
	TelegramChatId = "-1001683355327"
)

func BigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}
