package tools

import (
    "context"
    "encoding/json"
    "log"
    "MCP/internal/finam"
    orders "MCP/proto/grpc/tradeapi/v1/orders"
    trade "MCP/proto/grpc/tradeapi/v1"
    decimal "google.golang.org/genproto/googleapis/type/decimal"
)

type placeOrderArgs struct {
    AccountID string `json:"account_id"`
    Symbol    string `json:"symbol"`
    Quantity  string `json:"quantity"`
    Side      int32  `json:"side"` // trade.Side
    Type      int32  `json:"type"` // orders.OrderType
    TimeInForce int32 `json:"time_in_force"` // orders.TimeInForce
    LimitPrice string `json:"limit_price,omitempty"`
    StopPrice  string `json:"stop_price,omitempty"`
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

    ord := &orders.Order{
        AccountId:   args.AccountID,
        Symbol:      args.Symbol,
        Quantity:    &decimal.Decimal{Value: args.Quantity},
        Side:        trade.Side(args.Side),
        Type:        orders.OrderType(args.Type),
        TimeInForce: orders.TimeInForce(args.TimeInForce),
    }
    if args.LimitPrice != "" {
        ord.LimitPrice = &decimal.Decimal{Value: args.LimitPrice}
    }
    if args.StopPrice != "" {
        ord.StopPrice = &decimal.Decimal{Value: args.StopPrice}
    }
    resp, err := cli.PlaceOrder(ctx, ord)
	return resp, err
}
