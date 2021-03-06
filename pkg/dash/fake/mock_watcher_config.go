// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/pkg/dash (interfaces: WatcherConfig)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWatcherConfig is a mock of WatcherConfig interface.
type MockWatcherConfig struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherConfigMockRecorder
}

// MockWatcherConfigMockRecorder is the mock recorder for MockWatcherConfig.
type MockWatcherConfigMockRecorder struct {
	mock *MockWatcherConfig
}

// NewMockWatcherConfig creates a new mock instance.
func NewMockWatcherConfig(ctrl *gomock.Controller) *MockWatcherConfig {
	mock := &MockWatcherConfig{ctrl: ctrl}
	mock.recorder = &MockWatcherConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatcherConfig) EXPECT() *MockWatcherConfigMockRecorder {
	return m.recorder
}

// CurrentContext mocks base method.
func (m *MockWatcherConfig) CurrentContext() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentContext")
	ret0, _ := ret[0].(string)
	return ret0
}

// CurrentContext indicates an expected call of CurrentContext.
func (mr *MockWatcherConfigMockRecorder) CurrentContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentContext", reflect.TypeOf((*MockWatcherConfig)(nil).CurrentContext))
}

// UseContext mocks base method.
func (m *MockWatcherConfig) UseContext(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseContext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UseContext indicates an expected call of UseContext.
func (mr *MockWatcherConfigMockRecorder) UseContext(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseContext", reflect.TypeOf((*MockWatcherConfig)(nil).UseContext), arg0, arg1)
}

// UseFSContext mocks base method.
func (m *MockWatcherConfig) UseFSContext(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseFSContext", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UseFSContext indicates an expected call of UseFSContext.
func (mr *MockWatcherConfigMockRecorder) UseFSContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseFSContext", reflect.TypeOf((*MockWatcherConfig)(nil).UseFSContext), arg0)
}
