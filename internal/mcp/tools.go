package mcp

import (
    "encoding/json"
    "log"
    "context"
    tools "MCP/internal/tools"
    "MCP/internal/finam"
    marketdata "MCP/proto/grpc/tradeapi/v1/marketdata"
)

func init() {
    // Accounts:GetAccount
    tools.Register(tools.ToolDescriptor{
        Name: "accounts.getAccount",
        Description: "Get account details",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"account_id":{"type":"string"}},"required":["account_id"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ AccountID string `json:"account_id"` }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.GetAccount(ctx, in.AccountID)
    })

    // Accounts history
    tools.Register(tools.ToolDescriptor{
        Name: "accounts.trades",
        Description: "Get account trades",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"account_id":{"type":"string"},"limit":{"type":"integer"}},"required":["account_id"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ AccountID string; Limit int32 }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.Trades(ctx, in.AccountID, in.Limit)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "accounts.transactions",
        Description: "Get account transactions",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"account_id":{"type":"string"},"limit":{"type":"integer"}},"required":["account_id"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ AccountID string; Limit int32 }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.Transactions(ctx, in.AccountID, in.Limit)
    })

    // Assets endpoints
    tools.Register(tools.ToolDescriptor{
        Name: "assets.exchanges",
        Description: "List exchanges",
        InputSchema: json.RawMessage(`{"type":"object","properties":{}}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.Exchanges(ctx)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "assets.assets",
        Description: "List assets",
        InputSchema: json.RawMessage(`{"type":"object","properties":{}}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.AssetsList(ctx)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "assets.getAsset",
        Description: "Get asset details",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"symbol":{"type":"string"},"account_id":{"type":"string"}},"required":["symbol"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ Symbol, AccountID string }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.GetAsset(ctx, in.Symbol, in.AccountID)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "assets.getAssetParams",
        Description: "Get asset trading params",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"symbol":{"type":"string"},"account_id":{"type":"string"}},"required":["symbol","account_id"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ Symbol, AccountID string }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.GetAssetParams(ctx, in.Symbol, in.AccountID)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "assets.optionsChain",
        Description: "Get options chain",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"underlying_symbol":{"type":"string"}},"required":["underlying_symbol"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ UnderlyingSymbol string `json:"underlying_symbol"` }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.OptionsChain(ctx, in.UnderlyingSymbol)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "assets.schedule",
        Description: "Get trading schedule",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"symbol":{"type":"string"}},"required":["symbol"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ Symbol string }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.Schedule(ctx, in.Symbol)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "assets.clock",
        Description: "Get server clock",
        InputSchema: json.RawMessage(`{"type":"object","properties":{}}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.Clock(ctx)
    })

    // Marketdata unary methods
    tools.Register(tools.ToolDescriptor{
        Name: "marketdata.bars",
        Description: "Get historical bars",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"symbol":{"type":"string"},"timeframe":{"type":"integer"},"interval":{"type":"object"}},"required":["symbol","timeframe","interval"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{
            Symbol string `json:"symbol"`
            Timeframe int32 `json:"timeframe"`
            Interval struct{ StartTime string `json:"start_time"`; EndTime string `json:"end_time"` } `json:"interval"`
        }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        req := &marketdata.BarsRequest{ Symbol: in.Symbol, Timeframe: marketdata.TimeFrame(in.Timeframe) }
        return cli.Market.Bars(ctx, req)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "marketdata.lastQuote",
        Description: "Get last quote",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"symbol":{"type":"string"}},"required":["symbol"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ Symbol string }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.LastQuote(ctx, in.Symbol)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "marketdata.orderBook",
        Description: "Get order book",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"symbol":{"type":"string"}},"required":["symbol"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ Symbol string }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.OrderBook(ctx, in.Symbol)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "marketdata.latestTrades",
        Description: "Get latest trades",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"symbol":{"type":"string"}},"required":["symbol"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ Symbol string }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.LatestTrades(ctx, in.Symbol)
    })

    // Orders
    tools.Register(tools.ToolDescriptor{
        Name: "orders.placeOrder",
        Description: "Place an order",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"account_id":{"type":"string"},"symbol":{"type":"string"},"quantity":{"type":"string"},"side":{"type":"integer"},"type":{"type":"integer"},"time_in_force":{"type":"integer"},"limit_price":{"type":"string"},"stop_price":{"type":"string"}},"required":["account_id","symbol","quantity","side","type","time_in_force"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        return tools.Call(ctx, logger, "place_order", args)
    })

    // Auth
    tools.Register(tools.ToolDescriptor{
        Name: "auth.auth",
        Description: "Exchange API secret for JWT token",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"secret":{"type":"string"}},"required":["secret"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ Secret string }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.Authenticate(ctx, in.Secret)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "auth.tokenDetails",
        Description: "Get JWT token details",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"token":{"type":"string"}},"required":["token"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ Token string }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.TokenDetails(ctx, in.Token)
    })

    tools.Register(tools.ToolDescriptor{
        Name: "orders.cancelOrder",
        Description: "Cancel an order",
        InputSchema: json.RawMessage(`{"type":"object","properties":{"account_id":{"type":"string"},"order_id":{"type":"string"}},"required":["account_id","order_id"]}`),
    }, func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error) {
        var in struct{ AccountID, OrderID string }
        if err := json.Unmarshal(args, &in); err != nil { return nil, err }
        cli, err := finam.NewClient(); if err != nil { return nil, err }
        defer cli.Close()
        return cli.CancelOrder(ctx, in.AccountID, in.OrderID)
    })
}
