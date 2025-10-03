package main

import (
	finam "internal/finam"
	mcp "internal/mcp"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	logger := log.New(os.Stdout, "mcp-server ", log.LstdFlags|log.Lshortfile)

	wsHandler := mcp.NewServer(logger)

	http.HandleFunc("/ws", wsHandler.HandleWebSocket)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	logger.Printf("listening on %s (websocket at /ws)\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Fatalf("server error: %v", err)
	}
}
