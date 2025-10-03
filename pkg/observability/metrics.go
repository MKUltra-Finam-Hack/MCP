package observability

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ToolCallsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mcp_tool_calls_total",
			Help: "Total MCP tool calls",
		},
		[]string{"tool"},
	)
	ToolCallErrorsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mcp_tool_call_errors_total",
			Help: "Failed MCP tool calls",
		},
		[]string{"tool"},
	)
)
