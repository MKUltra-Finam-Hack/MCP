package finam

import (
	"context"
	"time"
)

type Order struct {
	ID        string    `json:"id"`
	Symbol    string    `json:"symbol"`
	Quantity  int       `json:"quantity"`
	Side      string    `json:"side"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Client) ListOrders(ctx context.Context, accountID string) ([]Order, error) {
	// TODO: Получить список ордеров
	return []Order{
		{ID: "1", Symbol: "SBER", Quantity: 10, Side: "BUY", Price: 150, Status: "ACTIVE", CreatedAt: time.Now()},
	}, nil
}

func (c *Client) CancelOrder(ctx context.Context, orderID string) (bool, error) {
	// TODO: реальный Finam API вызов отмены ордера
	return true, nil
}
