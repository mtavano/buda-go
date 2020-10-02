package buda

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mtavano/exchange"
	"github.com/pkg/errors"
)

const (
	statusReceived = "received"
	statusPending  = "pending" // Active

	statusTraded = "traded" // Executed

	statusCanceling = "canceling"
	statusCanceled  = "canceled" // Canceled
)

type Order struct {
	ID             uint64    `json:"id,omitempty"`
	Type           string    `json:"type,omitempty"`
	Status         string    `json:"state,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	MarketID       string    `json:"market_id,omitempty"`
	AccountID      uint64    `json:"account_id,omitempty"`
	FeeCurrency    string    `json:"fee_currency,omitempty"`
	PriceType      string    `json:"price_type,omitempty"`
	Limit          []string  `json:"limit,omitempty"`
	Amount         []string  `json:"amount,omitempty"`
	OriginalAmount []string  `json:"original_amount,omitempty"`
	TradedAmount   []string  `json:"traded_amount,omitempty"`
	TotalExchanged []string  `json:"total_exchanged,omitempty"`
	PaidFee        []string  `json:"paid_fee,omitempty"`
}

type postOrder struct {
	Type      string  `json:"type,omitempty"`
	PriceType string  `json:"price_type,omitempty"`
	Limit     float64 `json:"limit"`
	Amount    float64 `json:"amount,omitempty"`
}

func (b *Buda) CreateOrder(order *Order) (*Order, error) {
	var po postOrder
	pair := genericPairToBuda[order.Pair]

	body, err := b.MarshallBody(&po)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(ordersByMarektEndpoint, pair)
	res, err := b.makeRequest(http.MethodPost, url, body, true)
	if err != nil {
		return nil, err
	}

	payload := &struct {
		Order *Order `json:"order"`
	}{}
	err = b.scanBody(res, payload)
	fmt.Println(res.Body)
	if err != nil {
		return nil, err
	}

	return payload.Order, nil
}

func (b *Buda) GetOrder(id string) (*Order, error) {
	url := fmt.Sprintf(ordersByID, id)
	res, err := b.makeRequest(http.MethodGet, url, nil, true)
	if err != nil {
		return nil, errors.Wrap(err, "buda: b.GetOrder b.makeRequest error")
	}

	payload := &struct {
		Order *Order `json:"order"`
	}{}
	err = b.scanBody(res, payload)
	if err != nil {
		return nil, errors.Wrap(err, "buda: b.GetOrder b.scanBody error")
	}

	return payload.Order, nil
}

func (b *Buda) CancelOrder(id string) (*exchange.Order, error) {
	url := fmt.Sprintf(ordersByID, id)
	bodyRaw := &struct {
		State string `json:"state"`
	}{State: statusCanceling}

	body, err := b.MarshallBody(bodyRaw)
	if err != nil {
		return nil, errors.Wrap(err, "buda: b.CancelOrder b.Marshalllody error")
	}
	res, err := b.makeRequest(http.MethodPut, url, body, true)
	if err != nil {
		return nil, errors.Wrap(err, "buda: b.CancelOrder b.makeRequest error")
	}
	fmt.Println(res.StatusCode)

	payload := &struct {
		Order *Order `json:"order"`
	}{}
	err = b.scanBody(res, payload)
	if err != nil {
		return nil, errors.Wrap(err, "buda: b.CancelOrder b.scanBody error")
	}
	fmt.Println(payload)

	return payload.Order, nil
}
