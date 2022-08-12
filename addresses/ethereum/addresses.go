package ethereum_addresses

import (
	"go_defi/utils/constants"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	RPC_URL        = "wss://distinguished-holy-market.discover.quiknode.pro/8580eb52f1b8248ac369d3f4509c78016ad4b9e4/"
	MULTICALL_ADDR = common.HexToAddress("0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696")

	TOKENS = map[string]constants.Token{
		"dai": {
			Address:   common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
			Precision: big.NewInt(1e18),
			Size:      big.NewInt(1e18),
		},
		"usdc": {
			Address:   common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
			Precision: big.NewInt(1e6),
			Size:      big.NewInt(1e6),
		},
		"usdt": {
			Address:   common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
			Precision: big.NewInt(1e6),
			Size:      big.NewInt(1e6),
		},
		"wbtc": {
			Address:   common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"),
			Precision: big.NewInt(1e8),
			Size:      big.NewInt(0.00005e8),
		},
		"weth": {
			Address:   common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
			Precision: big.NewInt(1e18),
			Size:      big.NewInt(0.0005e18),
		},
	}

	LOOKUP = map[common.Address]string{}

	UNISWAP_V2_LIKE_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0xae461ca67b15dc8dc81ce7615e0320da1a9ab8d5"): {
			Name:     "dai-usdc",
			Tokens:   []constants.Token{TOKENS["dai"], TOKENS["usdc"]},
			Protocol: "UniswapV2",
		},
		common.HexToAddress("0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc"): {
			Name:     "usdc-weth",
			Tokens:   []constants.Token{TOKENS["usdc"], TOKENS["weth"]},
			Protocol: "UniswapV2",
		},
		common.HexToAddress("0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852"): {
			Name:     "weth-usdt",
			Tokens:   []constants.Token{TOKENS["weth"], TOKENS["usdt"]},
			Protocol: "UniswapV2",
		},
		common.HexToAddress("0xa478c2975ab1ea89e8196811f51a7b7ade33eb11"): {
			Name:     "dai-weth",
			Tokens:   []constants.Token{TOKENS["dai"], TOKENS["weth"]},
			Protocol: "UniswapV2",
		},
		common.HexToAddress("0x3041cbd36888becc7bbcbc0045e3b1f144466f5f"): {
			Name:     "usdc-usdt",
			Tokens:   []constants.Token{TOKENS["usdc"], TOKENS["usdt"]},
			Protocol: "UniswapV2",
		},
		common.HexToAddress("0xbb2b8038a1640196fbe3e38816f3e67cba72d940"): {
			Name:     "wbtc-weth",
			Tokens:   []constants.Token{TOKENS["wbtc"], TOKENS["weth"]},
			Protocol: "UniswapV2",
		},
		common.HexToAddress("0x397ff1542f962076d0bfe58ea045ffa2d347aca0"): {
			Name:     "usdc-weth",
			Tokens:   []constants.Token{TOKENS["usdc"], TOKENS["weth"]},
			Protocol: "SushiSwap",
		},
		common.HexToAddress("0x06da0fd433c1a5d7a4faa01111c044910a184553"): {
			Name:     "weth-usdt",
			Tokens:   []constants.Token{TOKENS["weth"], TOKENS["usdt"]},
			Protocol: "SushiSwap",
		},
		common.HexToAddress("0xceff51756c56ceffca006cd410b03ffc46dd3a58"): {
			Name:     "wbtc-weth",
			Tokens:   []constants.Token{TOKENS["wbtc"], TOKENS["weth"]},
			Protocol: "SushiSwap",
		},
		common.HexToAddress("0xc3d03e4f041fd4cd388c549ee2a29a9e5075882f"): {
			Name:     "dai-weth",
			Tokens:   []constants.Token{TOKENS["dai"], TOKENS["weth"]},
			Protocol: "SushiSwap",
		},
	}

	CURVE_STABLE_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0xbebc44782c7db0a1a60cb6fe97d0b483032ff1c7"): {
			Name:     "3pool",
			Tokens:   []constants.Token{TOKENS["dai"], TOKENS["usdc"], TOKENS["usdt"]},
			Protocol: "Curve",
		},
	}

	CURVE_CRYPTO_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0xd51a44d3fae010294c616388b506acda1bfaae46"): {
			Name:     "tricrypto2",
			Tokens:   []constants.Token{TOKENS["usdt"], TOKENS["wbtc"], TOKENS["weth"]},
			Protocol: "Curve",
		},
	}

	ALL_POOLS = map[common.Address]constants.Pool{}
)

func init() {
	for k, v := range UNISWAP_V2_LIKE_POOLS {
		v.Implementation = "UniswapV2"
		ALL_POOLS[k] = v
	}

	for k, v := range CURVE_STABLE_POOLS {
		v.Implementation = "CurveStableSwap"
		ALL_POOLS[k] = v
	}

	for k, v := range CURVE_CRYPTO_POOLS {
		v.Implementation = "CurveCryptoSwap"
		ALL_POOLS[k] = v
	}

	for k, v := range TOKENS {
		LOOKUP[v.Address] = k
	}
}
