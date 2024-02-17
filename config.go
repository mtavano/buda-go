package buda

const (
	// precision of parsing
	bitsLen64 = 64

	// Market Pairs
	PairBTCCLP = "btc-clp"
	PairETHCLP = "eth-clp"

	// endpoints
	marketOrderBookEndpoint = "/markets/%s/order_book"
	marketTradesEndpoint    = "/markets/%s/trades"
	marketTickerEndpoint    = "/markets/%s/ticker"
	accountBalanceEndpoint  = "/balances"
	ordersByMarektEndpoint  = "/markets/%s/orders"
	ordersByIDEndpoint      = "/orders/%s"
)
