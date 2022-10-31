package binance_tokens

type Token struct {
	Base string
	Quote string
}

var (
	TOKENS = map[string]Token{
		"BTCUSDT": {
			Base: "btc",
			Quote: "usdt",
		},
		"ETHUSDT": {
			Base: "eth",
			Quote: "usdt",
		},
		"LINKUSDT": {
			Base: "link",
			Quote: "usdt",
		},
		"UNIUSDT":{
			Base: "uni",
			Quote: "usdt",
		},
		"SUSHIUSDT": {
			Base: "sushi",
			Quote: "usdt",
		},
		"MKRUSDT": {
			Base: "mkr",
			Quote: "usdt",
		},
		"CRVUSDT": {
			Base: "crv",
			Quote: "usdt",
		},
		"AAVEUSDT": {
			Base: "aave",
			Quote: "usdt",
		},
	}
)