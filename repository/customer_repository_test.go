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

func TestCustomerRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCustomerRepository(ctrl)
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
				customer := domain.Customer{CustomerID: "C001", Name: "Bambang"}
				repo.EXPECT().Save(ctx, customer).Return(customer, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Customer{CustomerID: "C001", Name: "Bambang"})
			},
			expect:    domain.Customer{CustomerID: "C001", Name: "Bambang"},
			expectErr: false,
		},
		{
			name: "Save Failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Customer{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Customer{Name: "Invalid"})
			},
			expect:    domain.Customer{},
			expectErr: true,
		},
		{
			name: "Update Success",
			mock: func() {
				customer := domain.Customer{CustomerID: "C001", Name: "Updated Name"}
				repo.EXPECT().Update(ctx, customer).Return(customer, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, domain.Customer{CustomerID: "C001", Name: "Updated Name"})
			},
			expect:    domain.Customer{CustomerID: "C001", Name: "Updated Name"},
			expectErr: false,
		},
		{
			name: "FindById Success",
			mock: func() {
				repo.EXPECT().FindById(ctx, "C001").Return(domain.Customer{CustomerID: "C001", Name: "Bambang"}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, "C001")
			},
			expect:    domain.Customer{CustomerID: "C001", Name: "Bambang"},
			expectErr: false,
		},
		{
			name: "FindById Not Found",
			mock: func() {
				repo.EXPECT().FindById(ctx, "999").Return(domain.Customer{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, "999")
			},
			expect:    domain.Customer{},
			expectErr: true,
		},
		{
			name: "FindAll Success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Customer{{CustomerID: "C001", Name: "Bambang"}}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect:    []domain.Customer{{CustomerID: "C001", Name: "Bambang"}},
			expectErr: false,
		},
		{
			name: "Delete Success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Customer{CustomerID: "C001"}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Customer{CustomerID: "C001"})
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