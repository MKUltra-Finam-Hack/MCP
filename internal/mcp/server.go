package mcp

import (
	"encoding/json"
	"fmt"
	"internal/tools"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader
	logger   *log.Logger
}

func NewServer(logger *log.Logger) *Server {
	return &Server{
		upgrader: websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
		logger:   logger,
	}
}

// JSON-RPC 2.0 minimal structs
type rpcRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      any             `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type rpcResponse struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      any       `json:"id"`
	Result  any       `json:"result,omitempty"`
	Error   *rpcError `json:"error,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// MCP-specific message payloads (simplified)
type toolsListResult struct {
	Tools []tools.ToolDescriptor `json:"tools"`
}

type toolsCallParams struct {
	Name      string          `json:"name"`
	Arguments json.RawMessage `json:"arguments"`
}

func (s *Server) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer conn.Close()

	for {
		var req rpcRequest
		if err := conn.ReadJSON(&req); err != nil {
			s.logger.Printf("read error: %v", err)
			return
		}

		resp := rpcResponse{JSONRPC: "2.0", ID: req.ID}

		switch req.Method {
		case "tools/list":
			resp.Result = toolsListResult{Tools: tools.List()}
		case "tools/call":
			var p toolsCallParams
			if err := json.Unmarshal(req.Params, &p); err != nil {
				resp.Error = &rpcError{Code: -32602, Message: "invalid params", Data: err.Error()}
				break
			}
			result, callErr := tools.Call(r.Context(), s.logger, p.Name, p.Arguments)
			if callErr != nil {
				resp.Error = &rpcError{Code: -32000, Message: "tool error", Data: callErr.Error()}
			} else {
				resp.Result = result
			}
		default:
			resp.Error = &rpcError{Code: -32601, Message: fmt.Sprintf("method %s not found", req.Method)}
		}

		if err := conn.WriteJSON(resp); err != nil {
			s.logger.Printf("write error: %v", err)
			return
		}
	}
}
