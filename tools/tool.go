package tools

import (
	"context"
	"encoding/json"
)

//go:generate mockgen -source=tool.go -destination=mock_tool.go -package=tools
type tool interface {
	Name() string
	Description() string
	Invoke(ctx context.Context, input json.RawMessage) (json.RawMessage, error)
	InputArgsSchema() json.RawMessage
}
