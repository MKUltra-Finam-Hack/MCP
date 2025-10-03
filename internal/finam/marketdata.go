package finam

import (
	"context"
)

func (c *Client) GetQuote(ctx context.Context, symbol string) (*fproto.GetQuoteResponse, error) {
	return c.Market.GetQuote(ctx, &fproto.GetQuoteRequest{Symbol: symbol})
}
