package fantom_addresses

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	uniswap_query_addr      = "0xcc686630a9a1535d736b5ce05f8177cf09c9aab2"
	spookyswap_factory_addr = "0x152eE697f2E276fA89E96742e9bB9aB1F2E61bE3"
	spiritswap_factory_addr = "0xEF45d134b73241eDa7703fa787148D9C9F4950b0"
	dummy_addr              = "0xDC8d069681e2323820EDFb591aca0B5AbE783ca5"
)

var (
	RPC_URL = "wss://rpc.ankr.com/fantom/ws/09c5480d973de73d56110c6e85140402c805f6aec6b7380de71732e6e74eb16f"

	UNISWAP_QUERY_ADDR      = common.HexToAddress(uniswap_query_addr)
	SPOOKYSWAP_FACTORY_ADDR = common.HexToAddress(spookyswap_factory_addr)
	SPIRITSWAP_FACTORY_ADDR = common.HexToAddress(spiritswap_factory_addr)
	FACTORY_ADDRESSES       = []common.Address{SPOOKYSWAP_FACTORY_ADDR, SPIRITSWAP_FACTORY_ADDR}
	DUMMY_ADDR              = common.HexToAddress(dummy_addr)

	TOKEN_ADDRS = map[string]common.Address{
		"wftm": common.HexToAddress("0x21be370d5312f44cb42ce377bc9b8a0cef1a4c83"),
		"eth":  common.HexToAddress("0x74b23882a30290451A17c44f4F05243b6b58C76d"),
		"usdt": common.HexToAddress("0x049d68029688eabf473097a2fc38ef61633a3c7a"),
		"usdc": common.HexToAddress("0x04068da6c83afcfa0e13ba15a6696662335d5b75"),
		"dai":  common.HexToAddress("0x8d11ec38a3eb5e956b052f67da8bdc9bef8abf3e"),
		"btc":  common.HexToAddress("0x321162Cd933E2Be498Cd2267a90534A804051b11"),
	}

	TOKEN_DECIMALS = map[string]*big.Int{
		"wftm": big.NewInt(18),
		"eth":  big.NewInt(18),
		"usdt": big.NewInt(6),
		"usdc": big.NewInt(6),
		"dai":  big.NewInt(18),
		"btc":  big.NewInt(8),
	}

	REVERSE_NAMING = map[common.Address]string{}

	TRADABLE_TOKENS = []common.Address{
		TOKEN_ADDRS["wftm"],
		TOKEN_ADDRS["eth"],
		TOKEN_ADDRS["usdt"],
		TOKEN_ADDRS["usdc"],
		TOKEN_ADDRS["dai"],
		TOKEN_ADDRS["btc"],
	}
)

func init() {
	for key, value := range TOKEN_ADDRS {
		REVERSE_NAMING[value] = key
	}
}
