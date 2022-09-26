// Code generated by MockGen. DO NOT EDIT.
// Source: runner.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	testing "testing"

	gomock "github.com/golang/mock/gomock"
)

// MockRunner is a mock of Runner interface.
type MockRunner struct {
	ctrl     *gomock.Controller
	recorder *MockRunnerMockRecorder
}

// MockRunnerMockRecorder is the mock recorder for MockRunner.
type MockRunnerMockRecorder struct {
	mock *MockRunner
}

// NewMockRunner creates a new mock instance.
func NewMockRunner(ctrl *gomock.Controller) *MockRunner {
	mock := &MockRunner{ctrl: ctrl}
	mock.recorder = &MockRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunner) EXPECT() *MockRunnerMockRecorder {
	return m.recorder
}

// Install mocks base method.
func (m *MockRunner) Install(tb testing.TB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", tb)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockRunnerMockRecorder) Install(tb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockRunner)(nil).Install), tb)
}

// Uninstall mocks base method.
func (m *MockRunner) Uninstall(tb testing.TB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", tb)
	ret0, _ := ret[0].(error)
	return ret0
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockRunnerMockRecorder) Uninstall(tb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockRunner)(nil).Uninstall), tb)
}
