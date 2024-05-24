package func_call

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/invopop/jsonschema"
	"reflect"
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
func NewGenericFuncCallByPlainFunc(name, desc string, f interface{}) (*GenericFuncCall, error) {
	funcCall := &GenericFuncCall{
		name:        name,
		description: desc,
	}

	// valid f is a func
	if reflect.TypeOf(f).Kind() != reflect.Func {
		return nil, fmt.Errorf("f should be a func")
	}

	// valid f has 2 input params
	if reflect.TypeOf(f).NumIn() != 2 {
		return nil, fmt.Errorf("f should have 2 input params")
	}

	// valid f has 2 output params
	if reflect.TypeOf(f).NumOut() != 2 {
		return nil, fmt.Errorf("f should have 2 output params")
	}

	// valid f has 1st input param is context.Context
	if reflect.TypeOf(f).In(0) != reflect.TypeOf(context.Background()) {
		return nil, fmt.Errorf("f should have 1st input param is context.Context")
	}

	funcCall.fn = func(ctx context.Context, input json.RawMessage) (json.RawMessage, error) {
		req := reflect.New(reflect.TypeOf(f).In(1)).Interface()
		err := json.Unmarshal(input, req)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal input: %w", err)
		}
		ret := reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(req).Elem()})
		if ret[1].IsNil() {
			return json.Marshal(ret[0].Interface())
		}
		return nil, ret[1].Interface().(error)
	}

	// generate json schema
	schema := jsonschema.Reflect(reflect.New(reflect.TypeOf(f).In(1)).Interface())
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
