// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kamilsk/guard/pkg/transport/grpc/rpc (interfaces: Maintenance)

// Package rpc_test is a generated GoMock package.
package rpc_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	request "github.com/kamilsk/guard/pkg/service/types/request"
	response "github.com/kamilsk/guard/pkg/service/types/response"
	reflect "reflect"
)

// MockMaintenance is a mock of Maintenance interface
type MockMaintenance struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceMockRecorder
}

// MockMaintenanceMockRecorder is the mock recorder for MockMaintenance
type MockMaintenanceMockRecorder struct {
	mock *MockMaintenance
}

// NewMockMaintenance creates a new mock instance
func NewMockMaintenance(ctrl *gomock.Controller) *MockMaintenance {
	mock := &MockMaintenance{ctrl: ctrl}
	mock.recorder = &MockMaintenanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMaintenance) EXPECT() *MockMaintenanceMockRecorder {
	return m.recorder
}

// Install mocks base method
func (m *MockMaintenance) Install(arg0 context.Context, arg1 request.Install) response.Install {
	ret := m.ctrl.Call(m, "Install", arg0, arg1)
	ret0, _ := ret[0].(response.Install)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockMaintenanceMockRecorder) Install(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockMaintenance)(nil).Install), arg0, arg1)
}
