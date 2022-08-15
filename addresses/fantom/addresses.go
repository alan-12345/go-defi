package fantom_addresses

import (
	"go_defi/utils/constants"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	RPC_URL = "wss://rpc.ankr.com/fantom/ws/09c5480d973de73d56110c6e85140402c805f6aec6b7380de71732e6e74eb16f"
	
	COMPOUND_LIKE_PROTOCOLS = map[string]constants.Compound{
		"scream": {
			Unitroller: common.HexToAddress("0x260E596DAbE3AFc463e75B6CC05d8c46aCAcFB09"),
			StartBlock: big.NewInt(12149629),
		},
	}
)