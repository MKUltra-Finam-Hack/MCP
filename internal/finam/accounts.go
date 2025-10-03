package finam

import (
	"context"
	fproto "github.com/local/mcp-server/proto/finam"
)

func (c *Client) GetAccount(ctx context.Context, accountID string) (*fproto.GetAccountResponse, error) {
	req := &fproto.GetAccountRequest{AccountId: accountID}
	return c.Account.GetAccount(ctx, req)
}
