package finam

import (
	"google.golang.org/grpc"
	"log"
	"os"

	fproto "github.com/local/mcp-server/proto/finam"
)

type Client struct {
	Account fproto.AccountServiceClient
	Market  fproto.MarketDataServiceClient
	Order   fproto.OrderServiceClient
	conn    *grpc.ClientConn
}

func NewClient() (*Client, error) {
	addr := os.Getenv("FINAM_GRPC_ADDR")
	if addr == "" {
		addr = "localhost:9090"
	}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{
		Account: fproto.NewAccountServiceClient(conn),
		Market:  fproto.NewMarketDataServiceClient(conn),
		Order:   fproto.NewOrderServiceClient(conn),
		conn:    conn,
	}, nil
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
