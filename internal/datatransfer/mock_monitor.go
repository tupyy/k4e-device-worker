// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-flotta/flotta-device-worker/internal/datatransfer (interfaces: FileSync)

// Package datatransfer is a generated GoMock package.
package datatransfer

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileSync is a mock of FileSync interface.
type MockFileSync struct {
	ctrl     *gomock.Controller
	recorder *MockFileSyncMockRecorder
}

// MockFileSyncMockRecorder is the mock recorder for MockFileSync.
type MockFileSyncMockRecorder struct {
	mock *MockFileSync
}

// NewMockFileSync creates a new mock instance.
func NewMockFileSync(ctrl *gomock.Controller) *MockFileSync {
	mock := &MockFileSync{ctrl: ctrl}
	mock.recorder = &MockFileSyncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileSync) EXPECT() *MockFileSyncMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockFileSync) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockFileSyncMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockFileSync)(nil).Connect))
}

// SyncPath mocks base method.
func (m *MockFileSync) SyncPath(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncPath", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncPath indicates an expected call of SyncPath.
func (mr *MockFileSyncMockRecorder) SyncPath(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPath", reflect.TypeOf((*MockFileSync)(nil).SyncPath), arg0, arg1)
}
