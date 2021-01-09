package buda

const (
	// precision of parsing
	bitsLen64 = 64

	// Market Pairs
	pairBTCCLP = "btc-clp"
	pairETHCLP = "eth-clp"

	// endpoints
	marketTickerEndpoint   = "/markets/%s/ticker"
	accountBalanceEndpoint = "/balances"
	ordersByMarektEndpoint = "/markets/%s/orders"
	ordersByID             = "/orders/%s"
)
