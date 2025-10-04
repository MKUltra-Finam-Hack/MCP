package mcp

import (
    "encoding/json"
    "fmt"
    tools "MCP/internal/tools"
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

// HandleHTTPJSONRPC provides MCP over HTTP:
// - POST /mcp with JSON-RPC 2.0 request in body -> JSON response
// - POST /mcp?stream=1 with Accept: text/event-stream -> SSE single-shot stream
func (s *Server) HandleHTTPJSONRPC(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    // Read one JSON-RPC request
    var req rpcRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Process
    resp := s.processRPC(r, req)

    // Streaming via SSE when requested
    stream := r.URL.Query().Get("stream") == "1" || r.Header.Get("Accept") == "text/event-stream"
    if stream {
        w.Header().Set("Content-Type", "text/event-stream")
        w.Header().Set("Cache-Control", "no-cache")
        w.Header().Set("Connection", "keep-alive")
        flusher, ok := w.(http.Flusher)
        if !ok {
            http.Error(w, "stream unsupported", http.StatusInternalServerError)
            return
        }
        // write one event with response
        writeSSE := func(event string, v any) {
            b, _ := json.Marshal(v)
            _, _ = w.Write([]byte("event: "+event+"\n"))
            _, _ = w.Write([]byte("data: "))
            _, _ = w.Write(b)
            _, _ = w.Write([]byte("\n\n"))
            flusher.Flush()
        }
        writeSSE("response", resp)
        writeSSE("done", map[string]any{"id": resp.ID})
        return
    }

    // Regular JSON response
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) processRPC(r *http.Request, req rpcRequest) rpcResponse {
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
    return resp
}
