package buda

import (
	"fmt"
	"net/http"

	"github.com/blue-factory/cryptobot/pkg/exchange"
)

// GetBalances ...
func (b *Buda) GetBalances() ([]*exchange.Balance, error) {
	balances := make([]*exchange.Balance, 0)

	resource := fmt.Sprintf(accountBalanceEndpoint)
	res, err := b.makeRequest(http.MethodGet, resource, nil, true)
	if err != nil {
		return balances, err
	}

	payload := &struct {
		Balances []*rawBalance `json:"balances"`
	}{}
	err = b.scanBody(res, payload)

	for _, b := range payload.Balances {
		balance := b.To()
		fmt.Println(balance)
		balances = append(balances, balance)
	}

	return balances, nil
}

type rawBalance struct {
	ID                    string   `json:"id"`
	Amount                []string `json:"amount"`
	AvailableAmount       []string `json:"available_amount"`
	FrozenAmount          []string `json:"frozen_amount"`
	PendingWithdrawAmount []string `json:"pending_withdraw_amount"`
	AccountID             int      `json:"account_id"`
}

func (rb *rawBalance) To() *exchange.Balance {
	return &exchange.Balance{
		Available: rb.AvailableAmount[0],
		Wallet:    rb.ID,
		Balance:   rb.Amount[0],
	}
}
