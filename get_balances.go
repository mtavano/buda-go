package buda

import (
	"fmt"
	"net/http"
)

// GetBalances ...
func (b *Buda) GetBalances() ([]*Balance, error) {
	resource := fmt.Sprintf(accountBalanceEndpoint)
	res, err := b.makeRequest(http.MethodGet, resource, nil, true)
	if err != nil {
		return nil, err
	}

	payload := &struct {
		Balances []*Balance `json:"balances"`
	}{}
	err = b.scanBody(res, payload)

	return payload.Balances, nil
}

type Balance struct {
	ID                    string   `json:"id"`
	Amount                []string `json:"amount"`
	AvailableAmount       []string `json:"available_amount"`
	FrozenAmount          []string `json:"frozen_amount"`
	PendingWithdrawAmount []string `json:"pending_withdraw_amount"`
	AccountID             int      `json:"account_id"`
}
