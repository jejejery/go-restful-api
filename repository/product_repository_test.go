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

func TestProductRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockProductRepository(ctrl)
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
				product := domain.Product{ProductID: 1, Name: "Night Lamp"}
				repo.EXPECT().Save(ctx, product).Return(product, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Product{ProductID: 1, Name: "Night Lamp"})
			},
			expect:    domain.Product{ProductID: 1, Name: "Night Lamp"},
			expectErr: false,
		},
		{
			name: "Save Failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Product{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Product{Name: "Invalid"})
			},
			expect:    domain.Product{},
			expectErr: true,
		},
		{
			name: "Update Success",
			mock: func() {
				product := domain.Product{ProductID: 1, Name: "Updated Name"}
				repo.EXPECT().Update(ctx, product).Return(product, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, domain.Product{ProductID: 1, Name: "Updated Name"})
			},
			expect:    domain.Product{ProductID: 1, Name: "Updated Name"},
			expectErr: false,
		},
		{
			name: "FindById Success",
			mock: func() {
				repo.EXPECT().FindById(ctx, 1).Return(domain.Product{ProductID: 1, Name: "Night Lamp"}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 1)
			},
			expect:    domain.Product{ProductID: 1, Name: "Night Lamp"},
			expectErr: false,
		},
		{
			name: "FindById Not Found",
			mock: func() {
				repo.EXPECT().FindById(ctx, 999).Return(domain.Product{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 999)
			},
			expect:    domain.Product{},
			expectErr: true,
		},
		{
			name: "FindAll Success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Product{{ProductID: 1, Name: "Night Lamp"}}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect:    []domain.Product{{ProductID: 1, Name: "Night Lamp"}},
			expectErr: false,
		},
		{
			name: "Delete Success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Product{ProductID: 1}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Product{ProductID: 1})
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