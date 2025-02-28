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

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		Validate:           validate,
	}
}

// Create Customer
func (service *CustomerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) (web.CustomerResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.CustomerResponse{}, err
	}

	customer := domain.Customer{Name: request.Name, Email: request.Email, Phone: request.Phone, Address: request.Address, LoyaltyPts: 0}
	savedCustomer, err := service.CustomerRepository.Save(ctx, customer)
	if err != nil {
		return web.CustomerResponse{}, err
	}

	return helper.ToCustomerResponse(savedCustomer), nil
}

// Update Customer
func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) (web.CustomerResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.CustomerResponse{}, err
	}

	customer, err := service.CustomerRepository.FindById(ctx, request.CustomerID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.CustomerResponse{}, exception.NewNotFoundError("Customer not found")
	} else if err != nil {
		return web.CustomerResponse{}, err
	}

	customer.Name = request.Name
	customer.Email = request.Email
	customer.Phone = request.Phone
	customer.Address = request.Address
	
	updatedCustomer, err := service.CustomerRepository.Update(ctx, customer)
	if err != nil {
		return web.CustomerResponse{}, err
	}

	return helper.ToCustomerResponse(updatedCustomer), nil
}

// Delete Customer
func (service *CustomerServiceImpl) Delete(ctx context.Context, customerId string) error {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return exception.NewNotFoundError("Customer not found")
	} else if err != nil {
		return err
	}

	return service.CustomerRepository.Delete(ctx, customer)
}

// Find Customer By ID
func (service *CustomerServiceImpl) FindById(ctx context.Context, customerId string) (web.CustomerResponse, error) {
	customer, err := service.CustomerRepository.FindById(ctx, customerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.CustomerResponse{}, exception.NewNotFoundError("Customer not found")
	} else if err != nil {
		return web.CustomerResponse{}, err
	}

	return helper.ToCustomerResponse(customer), nil
}

// Find All Categories
func (service *CustomerServiceImpl) FindAll(ctx context.Context) ([]web.CustomerResponse, error) {
	categories, err := service.CustomerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.ToCustomerResponses(categories), nil
}
