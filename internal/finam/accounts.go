package finam

import (
	"context"
)

func (c *Client) GetAccount(ctx context.Context, accountID string) (*fproto.GetAccountResponse, error) {
	req := &fproto.GetAccountRequest{AccountId: accountID}
	return c.Account.GetAccount(ctx, req)
}
