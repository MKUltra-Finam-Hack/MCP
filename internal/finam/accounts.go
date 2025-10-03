package finam

import (
	"context"
)

type Account struct {
	ID       string  `json:"id"`
	Balance  float64 `json:"balance"`
	Free     float64 `json:"free"`
	Blocked  float64 `json:"blocked"`
	Currency string  `json:"currency"`
}

func (c *Client) GetAccount(ctx context.Context, accountID string) (*Account, error) {
	// TODO: Реальный API-запрос
	return &Account{
		ID:       accountID,
		Balance:  100000.0,
		Free:     80000.0,
		Blocked:  20000.0,
		Currency: "RUB",
	}, nil
}
