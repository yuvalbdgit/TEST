// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/pkg/dash (interfaces: FileWatcher)

// Package fake is a generated GoMock package.
package fake

import (
	reflect "reflect"

	fsnotify "github.com/fsnotify/fsnotify"
	gomock "github.com/golang/mock/gomock"
)

// MockFileWatcher is a mock of FileWatcher interface.
type MockFileWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockFileWatcherMockRecorder
}

// MockFileWatcherMockRecorder is the mock recorder for MockFileWatcher.
type MockFileWatcherMockRecorder struct {
	mock *MockFileWatcher
}

// NewMockFileWatcher creates a new mock instance.
func NewMockFileWatcher(ctrl *gomock.Controller) *MockFileWatcher {
	mock := &MockFileWatcher{ctrl: ctrl}
	mock.recorder = &MockFileWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileWatcher) EXPECT() *MockFileWatcherMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockFileWatcher) Add(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockFileWatcherMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockFileWatcher)(nil).Add), arg0)
}

// Errors mocks base method.
func (m *MockFileWatcher) Errors() chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Errors")
	ret0, _ := ret[0].(chan error)
	return ret0
}

// Errors indicates an expected call of Errors.
func (mr *MockFileWatcherMockRecorder) Errors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errors", reflect.TypeOf((*MockFileWatcher)(nil).Errors))
}

// Events mocks base method.
func (m *MockFileWatcher) Events() chan fsnotify.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Events")
	ret0, _ := ret[0].(chan fsnotify.Event)
	return ret0
}

// Events indicates an expected call of Events.
func (mr *MockFileWatcherMockRecorder) Events() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockFileWatcher)(nil).Events))
}
