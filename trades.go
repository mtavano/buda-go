package buda

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Transaction struct {
	Timestamp time.Time
	Amount    float64
	Price     float64
	OrderType string
}

// Trades is the response of the buda api for market traids according to a specific pair
type Trades struct {
	LastTimestamp time.Time
	MarketID      string
	Entries       []*Transaction
}

type tradeResponse struct {
	LastTimestamp string          `json:"last_timestamp"`
	MarketID      string          `json:"market_id"`
	Entries       [][]interface{} `json:"entries"`
}

// GetTrades ...
func (b *Buda) GetTrades(pair string) (*Trades, error) {
	url := fmt.Sprintf(marketTrades, pair)
	res, err := b.makeRequest(http.MethodGet, url, nil, false)
	if err != nil {
		return nil, errors.Wrap(err, "buda: b.GetTrades b.makeRequest error")
	}

	payload := &struct {
		Trades *tradeResponse `json:"trades"`
	}{}
	err = b.scanBody(res, payload)
	if err != nil {
		return nil, errors.Wrap(err, "buda: b.GetTrades b.scanBody error")
	}

	lastTimestamp, err := strconv.ParseInt(payload.Trades.LastTimestamp, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "buda: b.GetTrades ParseInt last_timestamp error")
	}

	entries := [][]string{}
	for _, entry := range payload.Trades.Entries {
		entries = append(entries, []string{
			fmt.Sprintf("%v", entry[0]),
			fmt.Sprintf("%v", entry[1]),
			fmt.Sprintf("%v", entry[2]),
			fmt.Sprintf("%v", entry[3]),
		})
	}

	trades := &Trades{
		LastTimestamp: time.Unix(lastTimestamp, 0),
		MarketID:      payload.Trades.MarketID,
		Entries:       entries,
	}

	return trades, nil
}
