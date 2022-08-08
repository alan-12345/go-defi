package ethereum_addresses

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	router_mainnet_addr  = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	mainnet_addr         = ""
	arb_contract_addr    = ""
	factory_mainnet_addr = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
)

var (
	MY_MAINNET_ADDR      = common.HexToAddress(mainnet_addr)
	UNISWAP_ROUTER       = common.HexToAddress(router_mainnet_addr)
	ARB_CONTRACT_ADDR    = common.HexToAddress(arb_contract_addr)
	UNISWAP_FACTORY_ADDR = common.HexToAddress(factory_mainnet_addr)
	WETH_ADDR            = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
)