package tools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/local/mcp-server/internal/finam"
)

type placeOrderArgs struct {
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

func validatePlaceOrder(a *placeOrderArgs) error {
	if a.AccountID == "" || a.Symbol == "" || a.Quantity == "" || a.Side == "" || a.Type == "" || a.TimeInForce == "" {
		return errors.New("missing required fields")
	}
	if a.ClientOrderID != nil && len(*a.ClientOrderID) > 20 {
		return errors.New("client_order_id must be <= 20 chars")
	}
	if a.Comment != nil && len(*a.Comment) > 128 {
		return errors.New("comment must be <= 128 chars")
	}
	return nil
}

var placeOrderSchema = json.RawMessage([]byte(`{
  "type":"object",
  "required":["account_id","symbol","quantity","side","type","time_in_force"],
  "properties":{
    "account_id":{"type":"string"},
    "symbol":{"type":"string"},
    "quantity":{"type":"string"},
    "side":{"type":"string","enum":["SIDE_BUY","SIDE_SELL"]},
    "type":{"type":"string","enum":["ORDER_TYPE_MARKET","ORDER_TYPE_LIMIT","ORDER_TYPE_STOP","ORDER_TYPE_STOP_LIMIT"]},
    "time_in_force":{"type":"string","enum":["TIME_IN_FORCE_DAY","TIME_IN_FORCE_GTC"]},
    "limit_price":{"type":"string"},
    "stop_price":{"type":"string"},
    "stop_condition":{"type":"string","enum":["STOP_CONDITION_UNSPECIFIED","STOP_CONDITION_LAST","STOP_CONDITION_BID","STOP_CONDITION_ASK"]},
    "client_order_id":{"type":"string","maxLength":20},
    "comment":{"type":"string","maxLength":128}
  }
}`))

func init() {
	Register(ToolDescriptor{
		Name:        "place_order",
		Description: "Create an exchange order on Finam",
		InputSchema: placeOrderSchema,
	}, placeOrder)
}

func placeOrder(ctx context.Context, logger *log.Logger, argsRaw json.RawMessage) (any, error) {
	if os.Getenv("FINAM_ACCESS_TOKEN") == "" {
		return nil, errors.New("FINAM_ACCESS_TOKEN is not set")
	}
	var args placeOrderArgs
	if err := json.Unmarshal(argsRaw, &args); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}
	if err := validatePlaceOrder(&args); err != nil {
		return nil, err
	}

	// Normalize enums to upper case just in case
	args.Side = strings.ToUpper(args.Side)
	args.Type = strings.ToUpper(args.Type)
	args.TimeInForce = strings.ToUpper(args.TimeInForce)

	cli := finam.NewClient()
	resp, err := cli.PlaceOrder(ctx, finam.PlaceOrderRequest{
		AccountID:     args.AccountID,
		Symbol:        args.Symbol,
		Quantity:      args.Quantity,
		Side:          args.Side,
		Type:          args.Type,
		TimeInForce:   args.TimeInForce,
		LimitPrice:    args.LimitPrice,
		StopPrice:     args.StopPrice,
		StopCondition: args.StopCondition,
		ClientOrderID: args.ClientOrderID,
		Comment:       args.Comment,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
