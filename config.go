package buda

const (
	// precision of parsing
	bitsLen64 = 64

	// Market Pairs
	PairBTCCLP = "btc-clp"
	PairETHCLP = "eth-clp"

	// endpoints
	marketTrades           = "/markets/%s/trades"
	marketTickerEndpoint   = "/markets/%s/ticker"
	accountBalanceEndpoint = "/balances"
	ordersByMarektEndpoint = "/markets/%s/orders"
	ordersByID             = "/orders/%s"
)
