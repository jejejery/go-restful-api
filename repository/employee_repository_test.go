package repository


import (
	"context"
	"errors"
	"github.com/jejejery/go-restful-api/model/domain"
	"github.com/jejejery/go-restful-api/repository/mocks"
	gomock "go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmployeeRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEmployeeRepository(ctrl)
	ctx := context.Background()

	tests := []struct {
		name      string
		mock      func()
		method    func() (interface{}, error)
		expect    interface{}
		expectErr bool
	}{
		{
			name: "Save Success",
			mock: func() {
				employee := domain.Employee{EmployeeID: "C001", Name: "Bambang"}
				repo.EXPECT().Save(ctx, employee).Return(employee, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Employee{EmployeeID: "C001", Name: "Bambang"})
			},
			expect:    domain.Employee{EmployeeID: "C001", Name: "Bambang"},
			expectErr: false,
		},
		{
			name: "Save Failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Employee{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Employee{Name: "Invalid"})
			},
			expect:    domain.Employee{},
			expectErr: true,
		},
		{
			name: "Update Success",
			mock: func() {
				employee := domain.Employee{EmployeeID: "C001", Name: "Updated Name"}
				repo.EXPECT().Update(ctx, employee).Return(employee, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, domain.Employee{EmployeeID: "C001", Name: "Updated Name"})
			},
			expect:    domain.Employee{EmployeeID: "C001", Name: "Updated Name"},
			expectErr: false,
		},
		{
			name: "FindById Success",
			mock: func() {
				repo.EXPECT().FindById(ctx, "C001").Return(domain.Employee{EmployeeID: "C001", Name: "Bambang"}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, "C001")
			},
			expect:    domain.Employee{EmployeeID: "C001", Name: "Bambang"},
			expectErr: false,
		},
		{
			name: "FindById Not Found",
			mock: func() {
				repo.EXPECT().FindById(ctx, "999").Return(domain.Employee{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, "999")
			},
			expect:    domain.Employee{},
			expectErr: true,
		},
		{
			name: "FindAll Success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Employee{{EmployeeID: "C001", Name: "Bambang"}}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect:    []domain.Employee{{EmployeeID: "C001", Name: "Bambang"}},
			expectErr: false,
		},
		{
			name: "Delete Success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Employee{EmployeeID: "C001"}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Employee{EmployeeID: "C001"})
			},
			expect:    nil,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := tt.method()

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, result)
			}
		})
	}
}