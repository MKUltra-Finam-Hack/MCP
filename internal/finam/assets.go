package finam

import (
	"context"
)

type Asset struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	ISIN   string `json:"isin"`
	Board  string `json:"board"`
}

func (c *Client) GetAsset(ctx context.Context, symbol string) (*Asset, error) {
	// TODO: Реальный запрос к Finam API
	return &Asset{
		Symbol: symbol,
		Name:   "Сбербанк России",
		ISIN:   "RU0009029540",
		Board:  "TQBR",
	}, nil
}

func (c *Client) SearchAssets(ctx context.Context, query string) ([]Asset, error) {
	// TODO: Фильтрация результатов поиска
	return []Asset{
		{Symbol: "SBER", Name: "Сбербанк России", ISIN: "RU0009029540", Board: "TQBR"},
		{Symbol: "GAZP", Name: "Газпром", ISIN: "RU0007661625", Board: "TQBR"},
	}, nil
}
