package controller

import (
	"github.com/jejejery/go-restful-api/exception"
	"github.com/jejejery/go-restful-api/model/web"
	"github.com/jejejery/go-restful-api/service"
	"github.com/gofiber/fiber/v2"
)

type EmployeeControllerImpl struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(employeeService service.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		EmployeeService: employeeService,
	}
}

// Create Employee
func (controller *EmployeeControllerImpl) Create(c *fiber.Ctx) error {
	employeeCreateRequest := new(web.EmployeeCreateRequest)
	if err := c.BodyParser(employeeCreateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	employeeResponse, err := controller.EmployeeService.Create(c.Context(), *employeeCreateRequest)
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
		Data:   employeeResponse,
	})
}

// Update Employee
func (controller *EmployeeControllerImpl) Update(c *fiber.Ctx) error {
	employeeUpdateRequest := new(web.EmployeeUpdateRequest)
	if err := c.BodyParser(employeeUpdateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	id := c.Params("employeeId")
	
	employeeUpdateRequest.EmployeeID = id


	employeeResponse, err := controller.EmployeeService.Update(c.Context(), *employeeUpdateRequest)
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
		Data:   employeeResponse,
	})
}

// Delete Employee
func (controller *EmployeeControllerImpl) Delete(c *fiber.Ctx) error {
	id := c.Params("employeeId")

	
	if err := controller.EmployeeService.Delete(c.Context(), id); err != nil {
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

// Find Employee By ID
func (controller *EmployeeControllerImpl) FindById(c *fiber.Ctx) error {
	id := c.Params("employeeId")
	

	employeeResponse, err := controller.EmployeeService.FindById(c.Context(), id)
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
		Data:   employeeResponse,
	})
}

// Find All Categories
func (controller *EmployeeControllerImpl) FindAll(c *fiber.Ctx) error {
	employeeResponses, err := controller.EmployeeService.FindAll(c.Context())
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
		Data:   employeeResponses,
	})
}
