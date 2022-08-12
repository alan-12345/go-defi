package constants

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Address   common.Address
	Precision *big.Int
	Size      *big.Int
}

type Pool struct {
	Name           string
	Tokens         []Token
	Protocol       string
	Implementation string
}

var (
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	Inf         = new(big.Float).SetInf(false)
	NegInf      = new(big.Float).SetInf(true)
	One         = new(big.Float).SetFloat64(1)
	NegOne      = new(big.Float).SetFloat64(-1)
	Zero        = new(big.Float).SetFloat64(0)
	TenInt      = big.NewInt(10)
)

func PrintDashed() {
	fmt.Println(strings.Repeat("-", 75))
}
