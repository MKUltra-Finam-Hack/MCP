package finam

import (
	"context"
    orders "MCP/proto/grpc/tradeapi/v1/orders"
)

func (c *Client) PlaceOrder(ctx context.Context, req *orders.Order) (*orders.OrderState, error) {
    return c.Orders.PlaceOrder(ctx, req)
}

func (c *Client) CancelOrder(ctx context.Context, accountID string, orderID string) (*orders.OrderState, error) {
    return c.Orders.CancelOrder(ctx, &orders.CancelOrderRequest{AccountId: accountID, OrderId: orderID})
}
