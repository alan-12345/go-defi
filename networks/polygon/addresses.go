package polygon_addresses

import (
	"go_defi/utils/constants"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	RPC_URL        = "wss://blue-warmhearted-meme.matic.discover.quiknode.pro/1b100ebeb9f9a885b9cf7545dc79c5750d2bc8ff/"
	MULTICALL_ADDR = common.HexToAddress("0x275617327c958bD06b5D6b871E7f491D76113dd8")
	BUNDLER_ADDR   = common.HexToAddress("0xa9c951777ecc5938E75748BBDfCf9B850410506a")

	TOKENS = map[string]constants.Token{
		"dai": {
			Address:  common.HexToAddress("0x8f3cf7ad23cd3cadbd9735aff958023239c6a063"),
			Decimals: 18,
			Size:     big.NewInt(1e18),
		},
		"usdc": {
			Address:  common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174"),
			Decimals: 6,
			Size:     big.NewInt(1e6),
		},
		"usdt": {
			Address:  common.HexToAddress("0xc2132d05d31c914a87c6611c10748aeb04b58e8f"),
			Decimals: 6,
			Size:     big.NewInt(1e6),
		},
		"wbtc": {
			Address:  common.HexToAddress("0x1bfd67037b42cf73acf2047067bd4f2c47d9bfd6"),
			Decimals: 8,
			Size:     big.NewInt(0.00005e8),
		},
		"weth": {
			Address:  common.HexToAddress("0x7ceb23fd6bc0add59e62ac25578270cff1b9f619"),
			Decimals: 18,
			Size:     big.NewInt(0.0005e18),
		},
		"wmatic": {
			Address:  common.HexToAddress("0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270"),
			Decimals: 18,
			Size:     big.NewInt(1e18),
		},
	}

	LOOKUP = map[common.Address]string{}

	UNISWAP_V3_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0x45dda9cb7c25131df268515131f647d726f50608"): {
			Name:     "usdc-weth (0.05%)",
			Tokens:   []constants.Token{TOKENS["usdc"], TOKENS["weth"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x50eaedb835021e4a108b7290636d62e9765cc6d7"): {
			Name:     "wbtc-weth (0.05%)",
			Tokens:   []constants.Token{TOKENS["wbtc"], TOKENS["weth"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x167384319b41f7094e62f7506409eb38079abff8"): {
			Name:     "wmatic-weth (0.3%)",
			Tokens:   []constants.Token{TOKENS["wmatic"], TOKENS["weth"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x86f1d8390222a3691c28938ec7404a1661e618e0"): {
			Name:     "wmatic-weth (0.05%)",
			Tokens:   []constants.Token{TOKENS["wmatic"], TOKENS["weth"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x847b64f9d3a95e977d157866447a5c0a5dfa0ee5"): {
			Name:     "wbtc-usdc (0.3%)",
			Tokens:   []constants.Token{TOKENS["wbtc"], TOKENS["usdc"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0xa374094527e1673a86de625aa59517c5de346d32"): {
			Name:     "wmatic-usdc (0.05%)",
			Tokens:   []constants.Token{TOKENS["wmatic"], TOKENS["usdc"]},
			Protocol: "UniswapV3",
		},
		common.HexToAddress("0x0e44ceb592acfc5d3f09d996302eb4c499ff8c10"): {
			Name:     "usdc-weth (0.3%)",
			Tokens:   []constants.Token{TOKENS["usdc"], TOKENS["weth"]},
			Protocol: "UniswapV3",
		},
	}

	UNISWAP_V2_LIKE_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0x853ee4b2a13f8a742d64c8f088be7ba2131f670d"): {
			Name:     "usdc-weth",
			Tokens:   []constants.Token{TOKENS["usdc"], TOKENS["weth"]},
			Protocol: "QuickSwap",
		},
		common.HexToAddress("0x6e7a5fafcec6bb1e78bae2a1f0b612012bf14827"): {
			Name:     "wmatic-usdc",
			Tokens:   []constants.Token{TOKENS["wmatic"], TOKENS["usdc"]},
			Protocol: "QuickSwap",
		},
		common.HexToAddress("0xadbf1854e5883eb8aa7baf50705338739e558e5b"): {
			Name:     "wmatic-weth",
			Tokens:   []constants.Token{TOKENS["wmatic"], TOKENS["weth"]},
			Protocol: "QuickSwap",
		},
		common.HexToAddress("0xf6422b997c7f54d1c6a6e103bcb1499eea0a7046"): {
			Name:     "weth-usdt",
			Tokens:   []constants.Token{TOKENS["weth"], TOKENS["usdt"]},
			Protocol: "QuickSwap",
		},
		common.HexToAddress("0x604229c960e5cacf2aaeac8be68ac07ba9df81c3"): {
			Name:     "wmatic-usdt",
			Tokens:   []constants.Token{TOKENS["wmatic"], TOKENS["usdt"]},
			Protocol: "QuickSwap",
		},
		common.HexToAddress("0x4a35582a710e1f4b2030a3f826da20bfb6703c09"): {
			Name:     "weth-dai",
			Tokens:   []constants.Token{TOKENS["weth"], TOKENS["dai"]},
			Protocol: "QuickSwap",
		},
		common.HexToAddress("0xf04adbf75cdfc5ed26eea4bbbb991db002036bdd"): {
			Name:     "usdc-dai",
			Tokens:   []constants.Token{TOKENS["usdc"], TOKENS["dai"]},
			Protocol: "QuickSwap",
		},
	}

	CURVE_STABLE_BASE_POOLS = map[common.Address]constants.Pool{}

	CURVE_STABLE_UNDERLYING_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0x445fe580ef8d70ff569ab36e80c647af338db351"): {
			Name:     "aave",
			Tokens:   []constants.Token{TOKENS["dai"], TOKENS["usdc"], TOKENS["usdt"]},
			Protocol: "Curve",
		},
	}

	CURVE_CRYPTO_BASE_POOLS = map[common.Address]constants.Pool{}

	CURVE_CRYPTO_UNDERLYING_POOLS = map[common.Address]constants.Pool{
		common.HexToAddress("0x1d8b86e3d88cdb2d34688e87e72f388cb541b7c8"): {
			Name:     "atricrypto3",
			Tokens:   []constants.Token{TOKENS["dai"], TOKENS["usdc"], TOKENS["usdt"], TOKENS["wbtc"], TOKENS["weth"]},
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

	for k, v := range TOKENS {
		LOOKUP[v.Address] = k
	}
}
