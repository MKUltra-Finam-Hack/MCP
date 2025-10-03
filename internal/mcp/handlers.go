package mcp

import (
	"context"
	"encoding/json"
	"log"

	"github.com/local/mcp-server/internal/finam"
)

// Шаблон RPC handler через gRPC-клиент
func HandleGetAccountBalance(ctx context.Context, logger *log.Logger, params json.RawMessage) (any, error) {
	var input struct {
		AccountID string `json:"account_id"`
	}
	if err := json.Unmarshal(params, &input); err != nil {
		return nil, err
	}
	cli, err := finam.NewClient()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	return cli.GetAccount(ctx, input.AccountID)
}
