// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/internal/objectvisitor (interfaces: ObjectHandler)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// MockObjectHandler is a mock of ObjectHandler interface.
type MockObjectHandler struct {
	ctrl     *gomock.Controller
	recorder *MockObjectHandlerMockRecorder
}

// MockObjectHandlerMockRecorder is the mock recorder for MockObjectHandler.
type MockObjectHandlerMockRecorder struct {
	mock *MockObjectHandler
}

// NewMockObjectHandler creates a new mock instance.
func NewMockObjectHandler(ctrl *gomock.Controller) *MockObjectHandler {
	mock := &MockObjectHandler{ctrl: ctrl}
	mock.recorder = &MockObjectHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectHandler) EXPECT() *MockObjectHandlerMockRecorder {
	return m.recorder
}

// AddEdge mocks base method.
func (m *MockObjectHandler) AddEdge(arg0 context.Context, arg1, arg2 *unstructured.Unstructured, arg3 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEdge", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEdge indicates an expected call of AddEdge.
func (mr *MockObjectHandlerMockRecorder) AddEdge(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEdge", reflect.TypeOf((*MockObjectHandler)(nil).AddEdge), arg0, arg1, arg2, arg3)
}

// Process mocks base method.
func (m *MockObjectHandler) Process(arg0 context.Context, arg1 *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process.
func (mr *MockObjectHandlerMockRecorder) Process(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockObjectHandler)(nil).Process), arg0, arg1)
}

// SetLevel mocks base method.
func (m *MockObjectHandler) SetLevel(arg0 string, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLevel", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// SetLevel indicates an expected call of SetLevel.
func (mr *MockObjectHandlerMockRecorder) SetLevel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLevel", reflect.TypeOf((*MockObjectHandler)(nil).SetLevel), arg0, arg1)
}
