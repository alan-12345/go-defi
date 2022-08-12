package constants

import (
	"fmt"
	"math/big"
	"strings"

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
	Precision *big.Int
	Size      *big.Int
}

type Pool struct {
	Name           string
	Tokens         []Token
	Fee            *big.Float
	Protocol       string
	Implementation string
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
)

func PrintDashed() {
	fmt.Println(strings.Repeat("-", 75))
}
