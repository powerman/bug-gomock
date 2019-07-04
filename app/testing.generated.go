// Code generated by MockGen. DO NOT EDIT.
// Source: app.go

// Package app is a generated GoMock package.
package app

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockApp is a mock of App interface
type MockApp struct {
	ctrl     *gomock.Controller
	recorder *MockAppMockRecorder
}

// MockAppMockRecorder is the mock recorder for MockApp
type MockAppMockRecorder struct {
	mock *MockApp
}

// NewMockApp creates a new mock instance
func NewMockApp(ctrl *gomock.Controller) *MockApp {
	mock := &MockApp{ctrl: ctrl}
	mock.recorder = &MockAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApp) EXPECT() *MockAppMockRecorder {
	return m.recorder
}

// Noop mocks base method
func (m *MockApp) Noop(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Noop", ctx)
}

// Noop indicates an expected call of Noop
func (mr *MockAppMockRecorder) Noop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Noop", reflect.TypeOf((*MockApp)(nil).Noop), ctx)
}