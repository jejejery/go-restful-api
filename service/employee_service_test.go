package service

import (
	"context"
	"errors"
	"github.com/jejejery/go-restful-api/model/domain"
	"github.com/jejejery/go-restful-api/model/web"
	"github.com/jejejery/go-restful-api/repository/mocks"
	"github.com/go-playground/validator/v10"
	"go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	mockValidator := validator.New()
	employeeService := NewEmployeeService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.EmployeeCreateRequest
		mock      func()
		expect    web.EmployeeResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: web.EmployeeCreateRequest{EmployeeID: "E001", 
												Name: "Hasan",
												Role: "Assassin",
											 	Email: "hasan@yahoo.com",
												Phone: "+6282345678912",
												DateHired: func() time.Time {	
													// ✅ Gunakan MustParse agar selalu berhasil
													_time, err := time.Parse("2006-01-02", "2021-01-01")
													if err != nil {
														panic(err) // Jika gagal, hentikan test segera
													}
													return _time
												}(),
											},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Employee{EmployeeID: "E001", Name: "Hasan"}, nil)
			},
			expect:    web.EmployeeResponse{EmployeeID: "E001", Name: "Hasan"},
			expectErr: false,
		},
		{
			name:      "validation error",
			input:     web.EmployeeCreateRequest{},
			mock:      func() {},
			expect:    web.EmployeeResponse{},
			expectErr: true,
		},
		{
			name:  "repository error",
			input: web.EmployeeCreateRequest{EmployeeID: "E001", 
											 Name: "Hasan",
											 Role: "Assassin",
											 Email: "hasan@yahoo.com",
											 Phone: "+6282345678912",
											 DateHired: helperDatetime("2021-01-01"),
											 },
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Employee{}, errors.New("database error"))
			},
			expect:    web.EmployeeResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := employeeService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockEmployeeRepository(ctrl)
	employeeService := NewEmployeeService(mockRepo, validator.New())

	tests := []struct {
		name       string
		employeeId string
		mock       func()
		expectErr  bool
	}{
		{
			name:       "success",
			employeeId: "E001",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "E001").Return(domain.Employee{EmployeeID: "E001", 
				Name: "Hasan",
				 Email: "hasan@yahoo.com",
				Phone: "+6282345678912",
				Role: "Assassin",
				DateHired: helperDatetime("2021-01-01"),}, nil)
				mockRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
		{
			name:       "not found",
			employeeId: "E002",
			mock: func() {
				mockRepo.EXPECT().FindById(gomock.Any(), "E002").Return(domain.Employee{}, errors.New("not found"))
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := employeeService.Delete(context.Background(), tt.employeeId)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateEmployee(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockEmployeeRepo *mocks.MockEmployeeRepository)
		input   web.EmployeeUpdateRequest
		expects error
	}{
		{
			name: "Success",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindById(gomock.Any(), "E001").
					Return(domain.Employee{EmployeeID: "E001", 
					Name: "Hasan",
					 Email: "hasan@yahoo.com",
					Phone: "+6282345678912",
					Role: "Assassin",
					DateHired: helperDatetime("2021-01-01")}, nil)
				mockEmployeeRepo.EXPECT().Update(gomock.Any(), gomock.Any()).
					Return(domain.Employee{EmployeeID: "E001", 
					Name: "Hasan",
					 Email: "hasan@yahoo.com",
					Phone: "+6282345678912",
					Role: "Assassin",
					DateHired: helperDatetime("2021-01-01")}, nil)
			},
			input:   web.EmployeeUpdateRequest{EmployeeID: "E001", Name: "Hasan"},
			expects: nil,
		},
		{
			name: "Employee Not Found",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindById(gomock.Any(), "E001").
					Return(domain.Employee{}, errors.New("not found"))
			},
			input:   web.EmployeeUpdateRequest{EmployeeID: "E001", Name: "Hasan"},
			expects: errors.New("not found"),
		},
		{
			name: "Database Error on Update",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindById(gomock.Any(), "E001").
					Return(domain.Employee{EmployeeID: "E001", Name: "Old Name"}, nil)
				mockEmployeeRepo.EXPECT().Update(gomock.Any(), gomock.Any()).
					Return(domain.Employee{}, errors.New("database error"))
			},
			input:   web.EmployeeUpdateRequest{EmployeeID: "E001", Name: "Hasuunnn"},
			expects: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockEmployeeRepo := mocks.NewMockEmployeeRepository(ctrl)
			tt.mock(mockEmployeeRepo)

			service := NewEmployeeService(mockEmployeeRepo, validator.New())
			_, err := service.Update(context.Background(), tt.input)

			if tt.expects != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expects.Error()) // Alternatif untuk assert.ErrorContains
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFindAllEmployees(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockEmployeeRepo *mocks.MockEmployeeRepository)
		expects []web.EmployeeResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindAll(gomock.Any()).Return([]domain.Employee{{EmployeeID: "E001", Name: "Employee 1"}}, nil)
			},
			expects: []web.EmployeeResponse{{EmployeeID: "E001", Name: "Employee 1"}},
			err:     nil,
		},
		{
			name: "Database Error",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindAll(gomock.Any()).Return(nil, errors.New("database error"))
			},
			expects: nil,
			err:     errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockEmployeeRepo := mocks.NewMockEmployeeRepository(ctrl)
			tt.mock(mockEmployeeRepo)

			service := NewEmployeeService(mockEmployeeRepo, validator.New())
			result, err := service.FindAll(context.Background())
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestFindByIdEmployee(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mockEmployeeRepo *mocks.MockEmployeeRepository)
		input   string
		expects web.EmployeeResponse
		err     error
	}{
		{
			name: "Success",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindById(gomock.Any(), "E001").Return(domain.Employee{EmployeeID: "E001", Name: "Employee 1"}, nil)
			},
			input:   "E001",
			expects: web.EmployeeResponse{EmployeeID: "E001", Name: "Employee 1"},
			err:     nil,
		},
		{
			name: "Not Found",
			mock: func(mockEmployeeRepo *mocks.MockEmployeeRepository) {
				mockEmployeeRepo.EXPECT().FindById(gomock.Any(), "E001").Return(domain.Employee{}, errors.New("not found"))
			},
			input:   "E001",
			expects: web.EmployeeResponse{},
			err:     errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockEmployeeRepo := mocks.NewMockEmployeeRepository(ctrl)
			tt.mock(mockEmployeeRepo)

			service := NewEmployeeService(mockEmployeeRepo, validator.New())
			result, err := service.FindById(context.Background(), tt.input)
			assert.Equal(t, tt.expects, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func helperDatetime(date string) time.Time {	
	// ✅ Gunakan MustParse agar selalu berhasil
	_time, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err) // Jika gagal, hentikan test segera
	}
	return _time
}
