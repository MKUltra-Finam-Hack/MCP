package finam

import (
    "context"
    accounts "MCP/proto/grpc/tradeapi/v1/accounts"
)

func (c *Client) Trades(ctx context.Context, accountID string, limit int32) (*accounts.TradesResponse, error) {
    return c.Account.Trades(ctx, &accounts.TradesRequest{AccountId: accountID, Limit: limit})
}

func (c *Client) Transactions(ctx context.Context, accountID string, limit int32) (*accounts.TransactionsResponse, error) {
    return c.Account.Transactions(ctx, &accounts.TransactionsRequest{AccountId: accountID, Limit: limit})
}

