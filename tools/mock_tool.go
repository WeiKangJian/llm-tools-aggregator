// Code generated by MockGen. DO NOT EDIT.
// Source: tool.go

// Package tools is a generated GoMock package.
package tools

import (
	context "context"
	json "encoding/json"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mocktool is a mock of tool interface.
type Mocktool struct {
	ctrl     *gomock.Controller
	recorder *MocktoolMockRecorder
}

// MocktoolMockRecorder is the mock recorder for Mocktool.
type MocktoolMockRecorder struct {
	mock *Mocktool
}

// NewMocktool creates a new mock instance.
func NewMocktool(ctrl *gomock.Controller) *Mocktool {
	mock := &Mocktool{ctrl: ctrl}
	mock.recorder = &MocktoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocktool) EXPECT() *MocktoolMockRecorder {
	return m.recorder
}

// Description mocks base method.
func (m *Mocktool) Description() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Description")
	ret0, _ := ret[0].(string)
	return ret0
}

// Description indicates an expected call of Description.
func (mr *MocktoolMockRecorder) Description() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Description", reflect.TypeOf((*Mocktool)(nil).Description))
}

// InputArgsSchema mocks base method.
func (m *Mocktool) InputArgsSchema() json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InputArgsSchema")
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}

// InputArgsSchema indicates an expected call of InputArgsSchema.
func (mr *MocktoolMockRecorder) InputArgsSchema() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InputArgsSchema", reflect.TypeOf((*Mocktool)(nil).InputArgsSchema))
}

// Invoke mocks base method.
func (m *Mocktool) Invoke(ctx context.Context, input json.RawMessage) (json.RawMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invoke", ctx, input)
	ret0, _ := ret[0].(json.RawMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Invoke indicates an expected call of Invoke.
func (mr *MocktoolMockRecorder) Invoke(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invoke", reflect.TypeOf((*Mocktool)(nil).Invoke), ctx, input)
}

// Name mocks base method.
func (m *Mocktool) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MocktoolMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*Mocktool)(nil).Name))
}
