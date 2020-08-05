// Code generated by MockGen. DO NOT EDIT.
// Source: src/domain/repository/department.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "accounting/src/domain/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDepartment is a mock of Department interface
type MockDepartment struct {
	ctrl     *gomock.Controller
	recorder *MockDepartmentMockRecorder
}

// MockDepartmentMockRecorder is the mock recorder for MockDepartment
type MockDepartmentMockRecorder struct {
	mock *MockDepartment
}

// NewMockDepartment creates a new mock instance
func NewMockDepartment(ctrl *gomock.Controller) *MockDepartment {
	mock := &MockDepartment{ctrl: ctrl}
	mock.recorder = &MockDepartmentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDepartment) EXPECT() *MockDepartmentMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockDepartment) Save(arg0 model.Department) *model.Department {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(*model.Department)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockDepartmentMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDepartment)(nil).Save), arg0)
}