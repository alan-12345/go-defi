package polygon_addresses

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	uniswap_query_addr     = "0x6BecC50C02dEF1B4f5b0c758eCdD4449f1695a1B"
	quickswap_factory_addr = "0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32"
	sushi_factory_addr     = "0xc35DADB65012eC5796536bD9864eD8773aBc74C4"
	fraxswap_factory_addr  = "0xc2544A32872A91F4A553b404C6950e89De901fdb"
)

var (
	RPC_URL                = "wss://ws-matic-mainnet.chainstacklabs.com"
	UNISWAP_QUERY_ADDR     = common.HexToAddress(uniswap_query_addr)
	QUICKSWAP_FACTORY_ADDR = common.HexToAddress(quickswap_factory_addr)
	SUSHI_FACTORY_ADDR     = common.HexToAddress(sushi_factory_addr)
	FRAXSWAP_FACTORY_ADDR  = common.HexToAddress(fraxswap_factory_addr)
	FACTORY_ADDRESSES      = []common.Address{QUICKSWAP_FACTORY_ADDR, SUSHI_FACTORY_ADDR}

	TOKEN_ADDRS = map[string]common.Address{
		"weth":   common.HexToAddress("0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"),
		"usdt":   common.HexToAddress("0xc2132d05d31c914a87c6611c10748aeb04b58e8f"),
		"usdc":   common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174"),
		"dai":    common.HexToAddress("0x8f3cf7ad23cd3cadbd9735aff958023239c6a063"),
		"wmatic": common.HexToAddress("0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270"),
		"wbtc":   common.HexToAddress("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6"),
		"frax":   common.HexToAddress("0x45c32fa6df82ead1e2ef74d17b76547eddfaff89"),
	}

	TOKEN_DECIMALS = map[string]*big.Int{
		"weth":   big.NewInt(18),
		"usdt":   big.NewInt(6),
		"usdc":   big.NewInt(6),
		"dai":    big.NewInt(18),
		"wmatic": big.NewInt(18),
		"wbtc":   big.NewInt(8),
		"frax":   big.NewInt(18),
	}

	REVERSE_NAMING = map[common.Address]string{}

	TRADABLE_TOKENS = []common.Address{
		TOKEN_ADDRS["weth"],
		TOKEN_ADDRS["usdt"],
		TOKEN_ADDRS["usdc"],
		TOKEN_ADDRS["wmatic"],
		TOKEN_ADDRS["wbtc"],
		TOKEN_ADDRS["frax"],
		TOKEN_ADDRS["dai"],
	}
)

func init() {
	for key, value := range TOKEN_ADDRS {
		REVERSE_NAMING[value] = key
	}
}
