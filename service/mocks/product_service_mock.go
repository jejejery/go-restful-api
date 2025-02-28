// Code generated by MockGen. DO NOT EDIT.
// Source: service/product_service.go
//
// Generated by this command:
//
//	mockgen -source=service/product_service.go -destination=service/mocks/product_service_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	web "github.com/jejejery/go-restful-api/model/web"
	gomock "go.uber.org/mock/gomock"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
	isgomock struct{}
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProductService) Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(web.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProductServiceMockRecorder) Create(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductService)(nil).Create), ctx, request)
}

// Delete mocks base method.
func (m *MockProductService) Delete(ctx context.Context, productId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, productId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductServiceMockRecorder) Delete(ctx, productId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductService)(nil).Delete), ctx, productId)
}

// FindAll mocks base method.
func (m *MockProductService) FindAll(ctx context.Context) ([]web.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]web.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockProductServiceMockRecorder) FindAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockProductService)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockProductService) FindById(ctx context.Context, productId int) (web.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, productId)
	ret0, _ := ret[0].(web.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockProductServiceMockRecorder) FindById(ctx, productId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockProductService)(nil).FindById), ctx, productId)
}

// Update mocks base method.
func (m *MockProductService) Update(ctx context.Context, request web.ProductUpdateRequest) (web.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, request)
	ret0, _ := ret[0].(web.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockProductServiceMockRecorder) Update(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductService)(nil).Update), ctx, request)
}
