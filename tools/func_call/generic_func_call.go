package func_call

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/invopop/jsonschema"
)

type GenericFuncCall struct {
	name            string
	description     string
	fn              func(ctx context.Context, input json.RawMessage) (json.RawMessage, error)
	inputArgsSchema json.RawMessage
}

func (g GenericFuncCall) Name() string {
	return g.name
}

func (g GenericFuncCall) Description() string {
	return g.description
}

func (g GenericFuncCall) Invoke(ctx context.Context, input json.RawMessage) (json.RawMessage, error) {
	return g.fn(ctx, input)
}

func (g GenericFuncCall) InputArgsSchema() json.RawMessage {
	return g.inputArgsSchema
}

// NewGenericFuncCallByPlainFunc generate generic func by plain func
// input func should be a func(ctx context.Context, req ReqType) (rsp rspType, error)
func NewGenericFuncCallByPlainFunc[T any, S any](name, desc string, f func(ctx context.Context, req T) (S, error)) (*GenericFuncCall, error) {
	funcCall := &GenericFuncCall{
		name:        name,
		description: desc,
	}

	funcCall.fn = func(ctx context.Context, input json.RawMessage) (json.RawMessage, error) {
		var req T
		err := json.Unmarshal(input, &req)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal input: %w", err)
		}

		res, err := f(ctx, req)
		if err != nil {
			return nil, err
		}

		output, err := json.Marshal(res)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal output: %w", err)
		}

		return output, nil
	}

	// generate json schema
	schema := jsonschema.Reflect(new(T))
	inputArgsSchema, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json schema: %w", err)
	}
	funcCall.inputArgsSchema = inputArgsSchema

	return funcCall, nil
}

// NewGenericFuncCall creates a new GenericFuncCall
func NewGenericFuncCall(
	name, description string,
	fn func(ctx context.Context, input json.RawMessage) (json.RawMessage, error),
	inputArgsSchema json.RawMessage) *GenericFuncCall {
	return &GenericFuncCall{
		name:            name,
		description:     description,
		fn:              fn,
		inputArgsSchema: inputArgsSchema,
	}
}
