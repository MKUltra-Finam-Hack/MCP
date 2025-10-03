package tools

import (
	"context"
	"encoding/json"
	"internal/finam"
	"log"
	"proto/finam"
)

type placeOrderArgs struct {
	AccountID string  `json:"account_id"`
	Symbol    string  `json:"symbol"`
	Quantity  int32   `json:"quantity"`
	Side      string  `json:"side"` // "BUY"/"SELL"
	Price     float64 `json:"price"`
	Type      string  `json:"type"`
}

func placeOrder(ctx context.Context, logger *log.Logger, argsRaw json.RawMessage) (any, error) {
	var args placeOrderArgs
	if err := json.Unmarshal(argsRaw, &args); err != nil {
		return nil, err
	}
	cli, err := finam.NewClient()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	grpcReq := &fproto.PlaceOrderRequest{
		AccountId: args.AccountID,
		Symbol:    args.Symbol,
		Quantity:  args.Quantity,
		Side:      args.Side,
		Price:     args.Price,
		Type:      args.Type,
	}
	resp, err := cli.PlaceOrder(ctx, grpcReq)
	return resp, err
}
