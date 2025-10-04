package finam

import (
	"context"
    assets "MCP/proto/grpc/tradeapi/v1/assets"
)

func (c *Client) Exchanges(ctx context.Context) (*assets.ExchangesResponse, error) {
    return c.Assets.Exchanges(ctx, &assets.ExchangesRequest{})
}

func (c *Client) AssetsList(ctx context.Context) (*assets.AssetsResponse, error) {
    return c.Assets.Assets(ctx, &assets.AssetsRequest{})
}

func (c *Client) GetAsset(ctx context.Context, symbol string, accountID string) (*assets.GetAssetResponse, error) {
    return c.Assets.GetAsset(ctx, &assets.GetAssetRequest{Symbol: symbol, AccountId: accountID})
}

func (c *Client) GetAssetParams(ctx context.Context, symbol string, accountID string) (*assets.GetAssetParamsResponse, error) {
    return c.Assets.GetAssetParams(ctx, &assets.GetAssetParamsRequest{Symbol: symbol, AccountId: accountID})
}

func (c *Client) OptionsChain(ctx context.Context, underlyingSymbol string) (*assets.OptionsChainResponse, error) {
    return c.Assets.OptionsChain(ctx, &assets.OptionsChainRequest{UnderlyingSymbol: underlyingSymbol})
}

func (c *Client) Schedule(ctx context.Context, symbol string) (*assets.ScheduleResponse, error) {
    return c.Assets.Schedule(ctx, &assets.ScheduleRequest{Symbol: symbol})
}

func (c *Client) Clock(ctx context.Context) (*assets.ClockResponse, error) {
    return c.Assets.Clock(ctx, &assets.ClockRequest{})
}
