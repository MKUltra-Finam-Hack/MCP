package finam

import (
	"context"
	"time"
)

type Quote struct {
	Symbol    string    `json:"symbol"`
	Bid       float64   `json:"bid"`
	Ask       float64   `json:"ask"`
	Last      float64   `json:"last"`
	Volume    int64     `json:"volume"`
	Timestamp time.Time `json:"timestamp"`
}

func (c *Client) GetQuote(ctx context.Context, symbol string) (*Quote, error) {
	// TODO: Реальный вызов Finam API
	return &Quote{
		Symbol:    symbol,
		Bid:       150.0,
		Ask:       151.0,
		Last:      150.5,
		Volume:    60000,
		Timestamp: time.Now(),
	}, nil
}

type Candle struct {
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume int64     `json:"volume"`
	Time   time.Time `json:"time"`
}

func (c *Client) GetCandles(ctx context.Context, symbol string, timeframe string, from, to time.Time) ([]Candle, error) {
	// TODO: Реальный вызов Finam API
	return []Candle{
		{Open: 150.2, High: 150.9, Low: 149.8, Close: 150.5, Volume: 32000, Time: from},
		{Open: 150.5, High: 150.8, Low: 149.9, Close: 150.0, Volume: 28000, Time: to},
	}, nil
}
