package buda

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

type Transaction struct {
	Timestamp string
	Amount    float64
	Price     float64
	OrderType string
}

// Trades is the response of the buda api for market traids according to a specific pair
type Trades struct {
	LastTimestamp string
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

	entries := make([]*Transaction, 0)
	for _, entry := range payload.Trades.Entries {
		entryPrice := fmt.Sprintf("%v", entry[0])
		price, err := strconv.ParseFloat(entryPrice, bitsLen64)
		if err != nil {
			return nil, errors.Wrap(err, "buda: b.GetTrades strconv.ParseFloat price error")
		}

		entryAount := fmt.Sprintf("%v", entry[1])
		amount, err := strconv.ParseFloat(entryAount, bitsLen64)
		if err != nil {
			return nil, errors.Wrap(err, "buda: b.GetTrades strconv.ParseFloat amount error")
		}

		entries = append(entries, &Transaction{
			Timestamp: payload.Trades.LastTimestamp,
			Price:     price,
			Amount:    amount,
			OrderType: fmt.Sprintf("%v", entry[3]),
		})
	}

	trades := &Trades{
		LastTimestamp: payload.Trades.LastTimestamp,
		MarketID:      payload.Trades.MarketID,
		Entries:       entries,
	}

	return trades, nil
}
