package buda

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Ticker ...
type Ticker struct {
	LastPrice         []string `json:"last_price"`
	MaxBid            []string `json:"max_bid"`
	MinAsk            []string `json:"min_ask"`
	PriceVariation24H string   `json:"price_variation_24h"`
	PriceVariation7D  string   `json:"price_variation_7d"`
	Volume            []string `json:"volume"`
}

// GetTicker expects a valid market pairr
// >> btc-clp
// other examples:
// eth-clp
// lc-clp
func (b *Buda) GetTicker(pair string) (*Ticker, error) {
	if pair == "" {
		return nil, errors.New("market pair cannot be empty")
	}

	url := fmt.Sprintf(marketTickerEndpoint, pair)
	res, err := b.makeRequest(http.MethodGet, url, nil, false)
	if err != nil {
		return nil, err
	}

	payload := &struct {
		Ticker *Ticker `json:"ticker"`
	}{}

	err = b.scanBody(res, payload)
	if err != nil {
		return nil, err
	}

	return payload.Ticker, nil
}
