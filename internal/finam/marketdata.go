package finam

import (
	"context"
    marketdata "MCP/proto/grpc/tradeapi/v1/marketdata"
)

func (c *Client) Bars(ctx context.Context, symbol string, timeframe marketdata.TimeFrame, interval *marketdata.BarsRequest) (*marketdata.BarsResponse, error) {
    return c.Market.Bars(ctx, interval)
}

func (c *Client) LastQuote(ctx context.Context, symbol string) (*marketdata.QuoteResponse, error) {
    return c.Market.LastQuote(ctx, &marketdata.QuoteRequest{Symbol: symbol})
}

func (c *Client) OrderBook(ctx context.Context, symbol string) (*marketdata.OrderBookResponse, error) {
    return c.Market.OrderBook(ctx, &marketdata.OrderBookRequest{Symbol: symbol})
}

func (c *Client) LatestTrades(ctx context.Context, symbol string) (*marketdata.LatestTradesResponse, error) {
    return c.Market.LatestTrades(ctx, &marketdata.LatestTradesRequest{Symbol: symbol})
}
