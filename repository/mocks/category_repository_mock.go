// Code generated by MockGen. DO NOT EDIT.
// Source: repository/category_repository.go
//
// Generated by this command:
//
//	mockgen -source=repository/category_repository.go -destination=repository/mocks/category_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/jejejery/go-restful-api/model/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
	isgomock struct{}
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCategoryRepository) Delete(ctx context.Context, category domain.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCategoryRepositoryMockRecorder) Delete(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategoryRepository)(nil).Delete), ctx, category)
}

// FindAll mocks base method.
func (m *MockCategoryRepository) FindAll(ctx context.Context) ([]domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCategoryRepositoryMockRecorder) FindAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCategoryRepository)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockCategoryRepository) FindById(ctx context.Context, categoryId int) (domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, categoryId)
	ret0, _ := ret[0].(domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockCategoryRepositoryMockRecorder) FindById(ctx, categoryId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockCategoryRepository)(nil).FindById), ctx, categoryId)
}

// Save mocks base method.
func (m *MockCategoryRepository) Save(ctx context.Context, category domain.Category) (domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, category)
	ret0, _ := ret[0].(domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockCategoryRepositoryMockRecorder) Save(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCategoryRepository)(nil).Save), ctx, category)
}

// Update mocks base method.
func (m *MockCategoryRepository) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, category)
	ret0, _ := ret[0].(domain.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCategoryRepositoryMockRecorder) Update(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategoryRepository)(nil).Update), ctx, category)
}
