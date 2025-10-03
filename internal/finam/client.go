package finam

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Client struct {
	baseURL string
	client  *http.Client
	token   string
}

func NewClient() *Client {
	base := os.Getenv("FINAM_API_BASE_URL")
	if base == "" {
		base = "https://api.finam.ru"
	}
	return &Client{
		baseURL: base,
		client:  &http.Client{Timeout: 30 * time.Second},
		token:   os.Getenv("FINAM_ACCESS_TOKEN"),
	}
}

// PlaceOrderRequest models REST fields compatible with grpc schema names.
type PlaceOrderRequest struct {
	AccountID     string  `json:"account_id"`
	Symbol        string  `json:"symbol"`
	Quantity      string  `json:"quantity"`
	Side          string  `json:"side"`
	Type          string  `json:"type"`
	TimeInForce   string  `json:"time_in_force"`
	LimitPrice    *string `json:"limit_price,omitempty"`
	StopPrice     *string `json:"stop_price,omitempty"`
	StopCondition *string `json:"stop_condition,omitempty"`
	ClientOrderID *string `json:"client_order_id,omitempty"`
	Comment       *string `json:"comment,omitempty"`
}

type PlaceOrderResponse struct {
	OrderID string `json:"order_id"`
	ExecID  string `json:"exec_id"`
	Status  string `json:"status"`
}

func (c *Client) PlaceOrder(ctx context.Context, req PlaceOrderRequest) (PlaceOrderResponse, error) {
	var out PlaceOrderResponse
	url := fmt.Sprintf("%s/v1/accounts/%s/orders", c.baseURL, req.AccountID)
	body, err := json.Marshal(req)
	if err != nil {
		return out, err
	}
	reqHTTP, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return out, err
	}
	reqHTTP.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		reqHTTP.Header.Set("Authorization", "Bearer "+c.token)
	}
	resp, err := c.client.Do(reqHTTP)
	if err != nil {
		return out, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		var e map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&e)
		return out, fmt.Errorf("finam error: status=%d body=%v", resp.StatusCode, e)
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return out, err
	}
	return out, nil
}
