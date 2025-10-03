package finam

import (
	"context"
	fproto "github.com/local/mcp-server/proto/finam"
)

func (c *Client) SearchAssets(ctx context.Context, query string, limit int32) ([]*fproto.Asset, error) {
	resp, err := c.Assets.SearchAssets(ctx, &fproto.SearchAssetsRequest{Query: query, Limit: limit})
	if err != nil {
		return nil, err
	}
	return resp.Assets, nil
}

func (c *Client) GetAsset(ctx context.Context, symbol string) (*fproto.Asset, error) {
	resp, err := c.Assets.GetAsset(ctx, &fproto.GetAssetRequest{Symbol: symbol})
	if err != nil {
		return nil, err
	}
	return resp.Asset, nil
}
