package ethereum_addresses

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	uniswap_query_addr   = "0x1090358861add9597133f64c05d7f4f5802c78d1"
	uniswap_factory_addr = "0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f"
	sushi_factory_addr   = "0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac"
)

var (
	RPC_URL = "wss://distinguished-holy-market.discover.quiknode.pro/8580eb52f1b8248ac369d3f4509c78016ad4b9e4/"

	UNISWAP_QUERY_ADDR   = common.HexToAddress(uniswap_query_addr)
	UNISWAP_FACTORY_ADDR = common.HexToAddress(uniswap_factory_addr)
	SUSHI_FACTORY_ADDR   = common.HexToAddress(sushi_factory_addr)
	FACTORY_ADDRESSES    = []common.Address{UNISWAP_FACTORY_ADDR, SUSHI_FACTORY_ADDR}

	TOKEN_ADDRS = map[string]common.Address{
		"weth": common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
		"usdt": common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
		"usdc": common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
		"dai":  common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
		"wbtc": common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"),
		"busd": common.HexToAddress("0x4fabb145d64652a948d72533023f6e7a623c7c53"),
		"fxs":  common.HexToAddress("0x3432b6a60d23ca0dfca7761b7ab56459d9c964d0"),
		"uni":  common.HexToAddress("0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"),
	}

	TOKEN_DECIMALS = map[string]*big.Int{
		"weth": big.NewInt(18),
		"usdt": big.NewInt(6),
		"usdc": big.NewInt(6),
		"dai":  big.NewInt(18),
		"wbtc": big.NewInt(8),
		"busd": big.NewInt(18),
		"fxs":  big.NewInt(18),
		"uni":  big.NewInt(18),
	}

	REVERSE_NAMING = map[common.Address]string{}

	TRADABLE_TOKENS = []common.Address{
		TOKEN_ADDRS["weth"],
		TOKEN_ADDRS["usdt"],
		TOKEN_ADDRS["usdc"],
		TOKEN_ADDRS["dai"],
		TOKEN_ADDRS["wbtc"],
		TOKEN_ADDRS["busd"],
		TOKEN_ADDRS["fxs"],
		TOKEN_ADDRS["uni"],
	}
)

func init() {
	for key, value := range TOKEN_ADDRS {
		REVERSE_NAMING[value] = key
	}
}
