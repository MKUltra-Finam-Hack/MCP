package finam

import (
    "context"
    auth "MCP/proto/grpc/tradeapi/v1/auth"
)

func (c *Client) Authenticate(ctx context.Context, secret string) (*auth.AuthResponse, error) {
    return c.Auth.Auth(ctx, &auth.AuthRequest{Secret: secret})
}

func (c *Client) TokenDetails(ctx context.Context, token string) (*auth.TokenDetailsResponse, error) {
    return c.Auth.TokenDetails(ctx, &auth.TokenDetailsRequest{Token: token})
}

