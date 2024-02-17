package buda

import (
	"fmt"
	"net/http"
)

type Balance struct {
	ID                    string   `json:"id"`
	Amount                []string `json:"amount"`
	AvailableAmount       []string `json:"available_amount"`
	FrozenAmount          []string `json:"frozen_amount"`
	PendingWithdrawAmount []string `json:"pending_withdraw_amount"`
	AccountID             int      `json:"account_id"`
}

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
	if err != nil {
		return nil, err
	}

	return payload.Balances, nil
}

func (b *Buda) GetBalanceByCurrency(currency string) (*Balance, error) {
	resource := fmt.Sprintf("%s/%s", accountBalanceEndpoint, currency)
	res, err := b.makeRequest(http.MethodGet, resource, nil, true)
	if err != nil {
		return nil, err
	}

	payload := &struct {
		Balance *Balance `json:"balance"`
	}{}
	err = b.scanBody(res, payload)
	if err != nil {
		return nil, err
	}

	return payload.Balance, nil
}
