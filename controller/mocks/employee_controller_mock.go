// Code generated by MockGen. DO NOT EDIT.
// Source: controller/employee_controller.go
//
// Generated by this command:
//
//	mockgen -source=controller/employee_controller.go -destination=controller/mocks/employee_controller_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockEmployeeController is a mock of EmployeeController interface.
type MockEmployeeController struct {
	ctrl     *gomock.Controller
	recorder *MockEmployeeControllerMockRecorder
	isgomock struct{}
}

// MockEmployeeControllerMockRecorder is the mock recorder for MockEmployeeController.
type MockEmployeeControllerMockRecorder struct {
	mock *MockEmployeeController
}

// NewMockEmployeeController creates a new mock instance.
func NewMockEmployeeController(ctrl *gomock.Controller) *MockEmployeeController {
	mock := &MockEmployeeController{ctrl: ctrl}
	mock.recorder = &MockEmployeeControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmployeeController) EXPECT() *MockEmployeeControllerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockEmployeeController) Create(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockEmployeeControllerMockRecorder) Create(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEmployeeController)(nil).Create), c)
}

// Delete mocks base method.
func (m *MockEmployeeController) Delete(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockEmployeeControllerMockRecorder) Delete(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockEmployeeController)(nil).Delete), c)
}

// FindAll mocks base method.
func (m *MockEmployeeController) FindAll(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockEmployeeControllerMockRecorder) FindAll(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockEmployeeController)(nil).FindAll), c)
}

// FindById mocks base method.
func (m *MockEmployeeController) FindById(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindById indicates an expected call of FindById.
func (mr *MockEmployeeControllerMockRecorder) FindById(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockEmployeeController)(nil).FindById), c)
}

// Update mocks base method.
func (m *MockEmployeeController) Update(c *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockEmployeeControllerMockRecorder) Update(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockEmployeeController)(nil).Update), c)
}
