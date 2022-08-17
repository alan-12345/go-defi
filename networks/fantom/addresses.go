package fantom_addresses

import (
	"go_defi/utils/constants"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// RPC_URL = "wss://tame-radial-butterfly.fantom.discover.quiknode.pro/6a25451fffebacfd9700b5ebf91fadb31c42f12b/"
	// RPC_URL="https://rpcapi.fantom.network"
	RPC_URL          = "wss://rpc.ankr.com/fantom/ws/09c5480d973de73d56110c6e85140402c805f6aec6b7380de71732e6e74eb16f"
	MULTICALL_ADDR   = common.HexToAddress("0xD98e3dBE5950Ca8Ce5a4b59630a5652110403E5c")
	EVENT_FETCH_SIZE = 500000
)

var (
	COMPOUND_LIKE_PROTOCOLS = map[string]constants.LendingProtocol{
		// "scream": {
		// 	Address: common.HexToAddress("0x260E596DAbE3AFc463e75B6CC05d8c46aCAcFB09"),
		// 	StartBlock: big.NewInt(12149629),
		// },
		"scream-v2": {
			Address: common.HexToAddress("0x3d3094aec3b63c744b9fe56397d36be568faebdf"),
			StartBlock: big.NewInt(39190091),
		},
	}
)
