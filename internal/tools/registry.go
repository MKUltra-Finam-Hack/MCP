package tools

import (
	"context"
	"encoding/json"
	"errors"
	"log"
)

type ToolDescriptor struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	InputSchema json.RawMessage `json:"inputSchema"`
}

type ToolFunc func(ctx context.Context, logger *log.Logger, args json.RawMessage) (any, error)

var registry = map[string]struct {
	desc ToolDescriptor
	fn   ToolFunc
}{}

func Register(desc ToolDescriptor, fn ToolFunc) {
	registry[desc.Name] = struct {
		desc ToolDescriptor
		fn   ToolFunc
	}{desc: desc, fn: fn}
}

func List() []ToolDescriptor {
	res := make([]ToolDescriptor, 0, len(registry))
	for _, r := range registry {
		res = append(res, r.desc)
	}
	return res
}

func Call(ctx context.Context, logger *log.Logger, name string, args json.RawMessage) (any, error) {
	entry, ok := registry[name]
	if !ok {
		return nil, errors.New("unknown tool: " + name)
	}
	return entry.fn(ctx, logger, args)
}
