package finam

import (
	"context"
	fproto "github.com/local/mcp-server/proto/finam"
)

func (c *Client) GetQuote(ctx context.Context, symbol string) (*fproto.GetQuoteResponse, error) {
	return c.Market.GetQuote(ctx, &fproto.GetQuoteRequest{Symbol: symbol})
}
