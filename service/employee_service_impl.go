package service

import (
	"context"
	"errors"
	"github.com/jejejery/go-restful-api/exception"
	"github.com/jejejery/go-restful-api/helper"
	"github.com/jejejery/go-restful-api/model/domain"
	"github.com/jejejery/go-restful-api/model/web"
	"github.com/jejejery/go-restful-api/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type EmployeeServiceImpl struct {
	EmployeeRepository repository.EmployeeRepository
	Validate           *validator.Validate
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository, validate *validator.Validate) EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
		Validate:           validate,
	}
}

// Create Employee
func (service *EmployeeServiceImpl) Create(ctx context.Context, request web.EmployeeCreateRequest) (web.EmployeeResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.EmployeeResponse{}, err
	}

	employee := domain.Employee{Name: request.Name, Role: request.Role, Email: request.Email, Phone: request.Phone, DateHired: request.DateHired}
	savedEmployee, err := service.EmployeeRepository.Save(ctx, employee)
	if err != nil {
		return web.EmployeeResponse{}, err
	}

	return helper.ToEmployeeResponse(savedEmployee), nil
}

// Update Employee
func (service *EmployeeServiceImpl) Update(ctx context.Context, request web.EmployeeUpdateRequest) (web.EmployeeResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.EmployeeResponse{}, err
	}

	employee, err := service.EmployeeRepository.FindById(ctx, request.EmployeeID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.EmployeeResponse{}, exception.NewNotFoundError("Employee not found")
	} else if err != nil {
		return web.EmployeeResponse{}, err
	}

	employee.Name = request.Name
	employee.Email = request.Email
	employee.Role = request.Role
	employee.Phone = request.Phone
	employee.DateHired = request.DateHired
	
	updatedEmployee, err := service.EmployeeRepository.Update(ctx, employee)
	if err != nil {
		return web.EmployeeResponse{}, err
	}

	return helper.ToEmployeeResponse(updatedEmployee), nil
}

// Delete Employee
func (service *EmployeeServiceImpl) Delete(ctx context.Context, employeeId string) error {
	employee, err := service.EmployeeRepository.FindById(ctx, employeeId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return exception.NewNotFoundError("Employee not found")
	} else if err != nil {
		return err
	}

	return service.EmployeeRepository.Delete(ctx, employee)
}

// Find Employee By ID
func (service *EmployeeServiceImpl) FindById(ctx context.Context, employeeId string) (web.EmployeeResponse, error) {
	employee, err := service.EmployeeRepository.FindById(ctx, employeeId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.EmployeeResponse{}, exception.NewNotFoundError("Employee not found")
	} else if err != nil {
		return web.EmployeeResponse{}, err
	}

	return helper.ToEmployeeResponse(employee), nil
}

// Find All Categories
func (service *EmployeeServiceImpl) FindAll(ctx context.Context) ([]web.EmployeeResponse, error) {
	categories, err := service.EmployeeRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.ToEmployeeResponses(categories), nil
}
