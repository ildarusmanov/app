// Code generated by MockGen. DO NOT EDIT.
// Source: ./goup/application.go

// Package mock_goup is a generated GoMock package.
package mock_goup

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// MockConfigManager is a mock of ConfigManager interface
type MockConfigManager struct {
	ctrl     *gomock.Controller
	recorder *MockConfigManagerMockRecorder
}

// MockConfigManagerMockRecorder is the mock recorder for MockConfigManager
type MockConfigManagerMockRecorder struct {
	mock *MockConfigManager
}

// NewMockConfigManager creates a new mock instance
func NewMockConfigManager(ctrl *gomock.Controller) *MockConfigManager {
	mock := &MockConfigManager{ctrl: ctrl}
	mock.recorder = &MockConfigManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigManager) EXPECT() *MockConfigManagerMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockConfigManager) Set(key string, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set
func (mr *MockConfigManagerMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockConfigManager)(nil).Set), key, value)
}

// Get mocks base method
func (m *MockConfigManager) Get(key string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockConfigManagerMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfigManager)(nil).Get), key)
}

// MockDependenciesManager is a mock of DependenciesManager interface
type MockDependenciesManager struct {
	ctrl     *gomock.Controller
	recorder *MockDependenciesManagerMockRecorder
}

// MockDependenciesManagerMockRecorder is the mock recorder for MockDependenciesManager
type MockDependenciesManagerMockRecorder struct {
	mock *MockDependenciesManager
}

// NewMockDependenciesManager creates a new mock instance
func NewMockDependenciesManager(ctrl *gomock.Controller) *MockDependenciesManager {
	mock := &MockDependenciesManager{ctrl: ctrl}
	mock.recorder = &MockDependenciesManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDependenciesManager) EXPECT() *MockDependenciesManagerMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockDependenciesManager) Add(name string, service interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", name, service)
}

// Add indicates an expected call of Add
func (mr *MockDependenciesManagerMockRecorder) Add(name, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockDependenciesManager)(nil).Add), name, service)
}

// Get mocks base method
func (m *MockDependenciesManager) Get(name string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDependenciesManagerMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDependenciesManager)(nil).Get), name)
}