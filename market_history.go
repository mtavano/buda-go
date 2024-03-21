package buda

import (
	"fmt"
	"net/http"
	"time"
)

type MarketHistory struct {
	Open       []float64 `json:"o"`
	Close      []float64 `json:"c"`
	High       []float64 `json:"h"`
	Low        []float64 `json:"l"`
	Volumes    []float64 `json:"v"`
	Timestamps []int64   `json:"t"`
}

func (b *Buda) GetMarketHistory(symbol string, from, to time.Time) (*MarketHistory, error) {
	res, err := b.makeRequest(
		http.MethodGet,
		fmt.Sprintf(marketHistoryEndpoint, symbol, from.Unix(), to.Unix()),
		nil,
		true,
	)
	if err != nil {
		fmt.Println("here 1 err", err.Error())
		return nil, err
	}

	var payload MarketHistory
	err = b.scanBody(res, &payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
