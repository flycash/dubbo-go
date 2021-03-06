// Code generated by MockGen. DO NOT EDIT.
// Source: invoker.go

// Package mock is a generated GoMock package.
package mock

import (
	"reflect"
)

import (
	"github.com/golang/mock/gomock"
)

import (
	"github.com/apache/dubbo-go/common"
	"github.com/apache/dubbo-go/protocol"
)

// MockInvoker is a mock of Invoker interface
type MockInvoker struct {
	ctrl     *gomock.Controller
	recorder *MockInvokerMockRecorder
}

// MockInvokerMockRecorder is the mock recorder for MockInvoker
type MockInvokerMockRecorder struct {
	mock *MockInvoker
}

// NewMockInvoker creates a new mock instance
func NewMockInvoker(ctrl *gomock.Controller) *MockInvoker {
	mock := &MockInvoker{ctrl: ctrl}
	mock.recorder = &MockInvokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInvoker) EXPECT() *MockInvokerMockRecorder {
	return m.recorder
}

// GetUrl mocks base method
func (m *MockInvoker) GetUrl() common.URL {
	ret := m.ctrl.Call(m, "GetUrl")
	ret0, _ := ret[0].(common.URL)
	return ret0
}

// GetUrl indicates an expected call of GetUrl
func (mr *MockInvokerMockRecorder) GetUrl() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrl", reflect.TypeOf((*MockInvoker)(nil).GetUrl))
}

// IsAvailable mocks base method
func (m *MockInvoker) IsAvailable() bool {
	ret := m.ctrl.Call(m, "IsAvailable")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAvailable indicates an expected call of IsAvailable
func (mr *MockInvokerMockRecorder) IsAvailable() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAvailable", reflect.TypeOf((*MockInvoker)(nil).IsAvailable))
}

// Destroy mocks base method
func (m *MockInvoker) Destroy() {
	m.ctrl.Call(m, "Destroy")
}

// Destroy indicates an expected call of Destroy
func (mr *MockInvokerMockRecorder) Destroy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockInvoker)(nil).Destroy))
}

// Invoke mocks base method
func (m *MockInvoker) Invoke(arg0 protocol.Invocation) protocol.Result {
	ret := m.ctrl.Call(m, "Invoke", arg0)
	ret0, _ := ret[0].(protocol.Result)
	return ret0
}

// Invoke indicates an expected call of Invoke
func (mr *MockInvokerMockRecorder) Invoke(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invoke", reflect.TypeOf((*MockInvoker)(nil).Invoke), arg0)
}
