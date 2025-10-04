package finam

import (
    "context"
    "crypto/tls"
    "os"
    "log"

    auth "MCP/proto/grpc/tradeapi/v1/auth"
    accounts "MCP/proto/grpc/tradeapi/v1/accounts"
    assets "MCP/proto/grpc/tradeapi/v1/assets"
    marketdata "MCP/proto/grpc/tradeapi/v1/marketdata"
    orders "MCP/proto/grpc/tradeapi/v1/orders"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/metadata"
)

type Client struct {
    Auth    auth.AuthServiceClient
    Account accounts.AccountsServiceClient
    Assets  assets.AssetsServiceClient
    Market  marketdata.MarketDataServiceClient
    Orders  orders.OrdersServiceClient
    conn    *grpc.ClientConn
    token   string
}

func NewClient() (*Client, error) {
    addr := os.Getenv("FINAM_GRPC_ADDR")
    if addr == "" {
        addr = "api.finam.ru:443"
    }
    jwt := os.Getenv("FINAM_JWT")
    secret := os.Getenv("FINAM_SECRET")
    debug := os.Getenv("FINAM_DEBUG") != ""
    unary := grpc.WithUnaryInterceptor(func(
        ctx context.Context,
        method string,
        req, reply any,
        cc *grpc.ClientConn,
        invoker grpc.UnaryInvoker,
        opts ...grpc.CallOption,
    ) error {
        // Per docs, each request must include JWT in Authorization header.
        // If FINAM_SECRET provided, mint fresh JWT per call; else use FINAM_JWT.
        if !isAuthMethod(method) {
            if secret != "" {
                a := auth.NewAuthServiceClient(cc)
                resp, err := a.Auth(ctx, &auth.AuthRequest{Secret: secret})
                if err != nil {
                    if debug { log.Printf("finam: Auth failed for %s: %v", method, err) }
                } else if tok := resp.GetToken(); tok != "" {
                    if debug { log.Printf("finam: minted JWT for %s (len=%d, head=%s...)", method, len(tok), safeHead(tok)) }
                    ctx = metadata.AppendToOutgoingContext(ctx, "authorization", tok)
                } else if debug {
                    log.Printf("finam: empty JWT received for %s", method)
                }
            } else if jwt != "" {
                if debug { log.Printf("finam: using provided JWT for %s (len=%d, head=%s...)", method, len(jwt), safeHead(jwt)) }
                ctx = metadata.AppendToOutgoingContext(ctx, "authorization", jwt)
            } else if debug {
                log.Printf("finam: no FINAM_SECRET/FINAM_JWT set for %s — call will likely fail", method)
            }
        }
        return invoker(ctx, method, req, reply, cc, opts...)
    })
    tlsCfg := &tls.Config{MinVersion: tls.VersionTLS12}
    conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg)), unary)
    if err != nil {
        return nil, err
    }
    cli := &Client{
        Auth:    auth.NewAuthServiceClient(conn),
        Account: accounts.NewAccountsServiceClient(conn),
        Assets:  assets.NewAssetsServiceClient(conn),
        Market:  marketdata.NewMarketDataServiceClient(conn),
        Orders:  orders.NewOrdersServiceClient(conn),
        conn:    conn,
    }
    return cli, nil
}

func fetchJWT(a auth.AuthServiceClient, secret string) (string, error) {
    resp, err := a.Auth(context.Background(), &auth.AuthRequest{Secret: secret})
    if err != nil { return "", err }
    return resp.Token, nil
}

func subscribeRenewal(a auth.AuthServiceClient, secret string, cli *Client) {
    stream, err := a.SubscribeJwtRenewal(context.Background(), &auth.SubscribeJwtRenewalRequest{Secret: secret})
    if err != nil { return }
    for {
        msg, err := stream.Recv()
        if err != nil { return }
        if msg.GetToken() != "" { cli.token = msg.GetToken() }
    }
}

func (c *Client) Close() {
    if c.conn != nil {
        c.conn.Close()
    }
}

func isAuthMethod(fullMethod string) bool {
    // full method form: /grpc.tradeapi.v1.auth.AuthService/Method
    if len(fullMethod) == 0 {
        return false
    }
    return contains(fullMethod, "/grpc.tradeapi.v1.auth.AuthService/")
}

func contains(s, sub string) bool {
    if len(sub) == 0 {
        return true
    }
    for i := 0; i+len(sub) <= len(s); i++ {
        if s[i:i+len(sub)] == sub {
            return true
        }
    }
    return false
}

func safeHead(tok string) string {
    if len(tok) <= 8 { return tok }
    return tok[:8]
}
