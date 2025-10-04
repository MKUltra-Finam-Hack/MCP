package finam

import (
	"context"
    accounts "MCP/proto/grpc/tradeapi/v1/accounts"
)

func (c *Client) GetAccount(ctx context.Context, accountID string) (*accounts.GetAccountResponse, error) {
    req := &accounts.GetAccountRequest{AccountId: accountID}
    return c.Account.GetAccount(ctx, req)
}
