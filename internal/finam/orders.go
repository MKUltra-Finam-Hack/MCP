package finam

import (
	"context"
	fproto "github.com/local/mcp-server/proto/finam"
)

func (c *Client) PlaceOrder(ctx context.Context, req *fproto.PlaceOrderRequest) (*fproto.PlaceOrderResponse, error) {
	return c.Order.PlaceOrder(ctx, req)
}

func (c *Client) CancelOrder(ctx context.Context, orderID string) (*fproto.CancelOrderResponse, error) {
	return c.Order.CancelOrder(ctx, &fproto.CancelOrderRequest{OrderId: orderID})
}
