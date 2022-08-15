package ethereum_addresses

import (
	"go_defi/utils/constants"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	RPC_URL                = "wss://distinguished-holy-market.discover.quiknode.pro/8580eb52f1b8248ac369d3f4509c78016ad4b9e4/"
	MULTICALL_ADDR         = common.HexToAddress("0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696")
	BUNDLER_ADDR           = common.HexToAddress("0x9528e3AA3CD60030Bb34cB0b0539Ec0c0b6437c8")

	TOKENS = map[string]constants.Token{
		"dai": {
			Address:  common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
			Decimals: 18,
			Size:     big.NewInt(1e18),
		},
		"usdc": {
			Address:  common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
			Decimals: 6,
			Size:     big.NewInt(1e6),
		},
		"usdt": {
			Address:  common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
			Decimals: 6,
			Size:     big.NewInt(1e6),
		},
		"wbtc": {
			Address:  common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"),
			Decimals: 8,
			Size:     big.NewInt(0.00005e8),
		},
		"weth": {
			Address:  common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
			Decimals: 18,
			Size:     big.NewInt(0.0005e18),
		},
		"susd": {
			Address:  common.HexToAddress("0x57Ab1ec28D129707052df4dF418D58a2D46d5f51"),
			Decimals: 18,
			Size:     big.NewInt(1e18),
		},
		"frax": {
			Address:  common.HexToAddress("0x57Ab1ec28D129707052df4dF418D58a2D46d5f51"),
			Decimals: 18,
			Size:     big.NewInt(1e18),
		},
	}

	LOOKUP = map[common.Address]string{}

	UNISWAP_V3_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0x5777d92f208679db4b9778590fa3cab3ac9e2168"): {
			Name:     "dai-usdc (0.01%)",
			Tokens:   []constants.Token{TOKENS["dai"], TOKENS["usdc"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x6c6bc977e13df9b0de53b251522280bb72383700"): {
			Name:     "dai-usdc (0.05%)",
			Tokens:   []constants.Token{TOKENS["dai"], TOKENS["usdc"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x8ad599c3a0ff1de082011efddc58f1908eb6e6d8"): {
			Name:     "usdc-weth (0.3%)",
			Tokens:   []constants.Token{TOKENS["usdc"], TOKENS["weth"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"): {
			Name:     "usdc-weth (0.05%)",
			Tokens:   []constants.Token{TOKENS["usdc"], TOKENS["weth"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x3416cf6c708da44db2624d63ea0aaef7113527c6"): {
			Name:     "usdc-usdt (0.01%)",
			Tokens:   []constants.Token{TOKENS["usdc"], TOKENS["usdt"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x4585fe77225b41b697c938b018e2ac67ac5a20c0"): {
			Name:     "wbtc-weth (0.05%)",
			Tokens:   []constants.Token{TOKENS["wbtc"], TOKENS["weth"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x11b815efb8f581194ae79006d24e0d814b7697f6"): {
			Name:     "weth-usdt (0.05%)",
			Tokens:   []constants.Token{TOKENS["weth"], TOKENS["usdt"]},
			Protocol: "UniswapV3",
		},
	}

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

	CURVE_STABLE_BASE_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0xbebc44782c7db0a1a60cb6fe97d0b483032ff1c7"): {
			Name:     "3pool",
			Tokens:   []constants.Token{TOKENS["dai"], TOKENS["usdc"], TOKENS["usdt"]},
			Protocol: "Curve",
		},
	}

	CURVE_STABLE_UNDERLYING_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0xA2B47E3D5c44877cca798226B7B8118F9BFb7A56"): {
			Name:     "compound",
			Tokens:   []constants.Token{TOKENS["dai"], TOKENS["usdc"]},
			Protocol: "Curve",
		},
	}

	CURVE_CRYPTO_BASE_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0xd51a44d3fae010294c616388b506acda1bfaae46"): {
			Name:     "tricrypto2",
			Tokens:   []constants.Token{TOKENS["usdt"], TOKENS["wbtc"], TOKENS["weth"]},
			Protocol: "Curve",
		},
	}

	CURVE_CRYPTO_UNDERLYING_POOLS = map[common.Address]constants.Pool{}

	CURVE_META_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0xd632f22692fac7611d2aa1c0d552930d43caed3b"): {
			Name:     "frax",
			Tokens:   []constants.Token{TOKENS["frax"], TOKENS["dai"], TOKENS["usdc"], TOKENS["usdt"]},
			Protocol: "Curve",
		},
	}

	ALL_POOLS = map[common.Address]constants.Pool{}
)

func init() {
	for k, v := range UNISWAP_V2_LIKE_POOLS {
		v.SwapType = constants.SwapTypes["UniswapV2"]
		ALL_POOLS[k] = v
	}

	for k, v := range UNISWAP_V3_POOLS {
		v.SwapType = constants.SwapTypes["UniswapV3"]
		ALL_POOLS[k] = v
	}

	for k, v := range CURVE_STABLE_BASE_POOLS {
		v.SwapType = constants.SwapTypes["CurveStableBase"]
		ALL_POOLS[k] = v
	}

	for k, v := range CURVE_STABLE_UNDERLYING_POOLS {
		v.SwapType = constants.SwapTypes["CurveStableUnderlying"]
		ALL_POOLS[k] = v
	}

	for k, v := range CURVE_CRYPTO_BASE_POOLS {
		v.SwapType = constants.SwapTypes["CurveCryptoBase"]
		ALL_POOLS[k] = v
	}

	for k, v := range CURVE_CRYPTO_UNDERLYING_POOLS {
		v.SwapType = constants.SwapTypes["CurveCryptoUnderlying"]
		ALL_POOLS[k] = v
	}

	for k, v := range CURVE_META_POOLS {
		v.SwapType = constants.SwapTypes["CurveMetaPool"]
		ALL_POOLS[k] = v
	}

	for k, v := range TOKENS {
		LOOKUP[v.Address] = k
	}
}
