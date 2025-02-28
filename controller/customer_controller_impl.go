package controller

import (
	"github.com/jejejery/go-restful-api/exception"
	"github.com/jejejery/go-restful-api/model/web"
	"github.com/jejejery/go-restful-api/service"
	"github.com/gofiber/fiber/v2"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

// Create Customer
func (controller *CustomerControllerImpl) Create(c *fiber.Ctx) error {
	customerCreateRequest := new(web.CustomerCreateRequest)
	if err := c.BodyParser(customerCreateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	customerResponse, err := controller.CustomerService.Create(c.Context(), *customerCreateRequest)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "Created",
		Data:   customerResponse,
	})
}

// Update Customer
func (controller *CustomerControllerImpl) Update(c *fiber.Ctx) error {
	customerUpdateRequest := new(web.CustomerUpdateRequest)
	if err := c.BodyParser(customerUpdateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	id := c.Params("customerId")
	
	customerUpdateRequest.CustomerID = id


	customerResponse, err := controller.CustomerService.Update(c.Context(), *customerUpdateRequest)
	if err != nil {
		if _, ok := err.(exception.NotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "Not Found",
				Data:   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	})
}

// Delete Customer
func (controller *CustomerControllerImpl) Delete(c *fiber.Ctx) error {
	id := c.Params("customerId")

	
	if err := controller.CustomerService.Delete(c.Context(), id); err != nil {
		if _, ok := err.(exception.NotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "Not Found",
				Data:   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "Deleted Successfully",
	})
}

// Find Customer By ID
func (controller *CustomerControllerImpl) FindById(c *fiber.Ctx) error {
	id := c.Params("customerId")
	

	customerResponse, err := controller.CustomerService.FindById(c.Context(), id)
	if err != nil {
		if _, ok := err.(exception.NotFoundError); ok {
			return c.Status(fiber.StatusNotFound).JSON(web.WebResponse{
				Code:   fiber.StatusNotFound,
				Status: "Not Found",
				Data:   err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	})
}

// Find All Categories
func (controller *CustomerControllerImpl) FindAll(c *fiber.Ctx) error {
	customerResponses, err := controller.CustomerService.FindAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.WebResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   customerResponses,
	})
}
