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
	Address   common.Address
	Decimals  int64
	Precision *big.Int
	Size      *big.Int
}

type Pool struct {
	Name           string
	Tokens         []Token
	Fee            *big.Float
	Protocol       string
	Implementation string
	SwapType       uint8
}

var (
	ZeroAddress          = common.HexToAddress("0x0000000000000000000000000000000000000000")
	Inf                  = new(big.Float).SetInf(false)
	NegInf               = new(big.Float).SetInf(true)
	One                  = new(big.Float).SetFloat64(1)
	NegOne               = new(big.Float).SetFloat64(-1)
	Zero                 = new(big.Float).SetFloat64(0)
	TenInt               = big.NewInt(10)
	PointThreePercent    = big.NewFloat(0.003)
	PointZeroFivePercent = big.NewFloat(0.0005)
	PointZeroOnePercent  = big.NewFloat(0.0001)
	SwapTypes            = map[string]uint8{
		"UniswapV2":             0,
		"UniswapV3":             1,
		"CurveStableBase":       2,
		"CurveStableUnderlying": 3,
		"CurveCryptoBase":       4,
		"CurveCryptoUnderlying": 5,
	}
	ReallyBigInt = new(big.Int).Sub(BigPow(2, 128), big.NewInt(1))
)

func BigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}
