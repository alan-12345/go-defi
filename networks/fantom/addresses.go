package fantom_addresses

import (
	"go_defi/utils/constants"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	RPC_URL = "wss://tame-radial-butterfly.fantom.discover.quiknode.pro/6a25451fffebacfd9700b5ebf91fadb31c42f12b/"

	COMPOUND_LIKE_PROTOCOLS = map[string]constants.Compound{
		"scream": {
			Unitroller: common.HexToAddress("0x260E596DAbE3AFc463e75B6CC05d8c46aCAcFB09"),
			StartBlock: big.NewInt(12149629),
		},
		"scream-v2": {
			Unitroller: common.HexToAddress("0x3d3094aec3b63c744b9fe56397d36be568faebdf"),
			StartBlock: big.NewInt(39190091),
		},
	}
)
