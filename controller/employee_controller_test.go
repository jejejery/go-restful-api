package controller

import (
	"bytes"
	"encoding/json"
	"github.com/jejejery/go-restful-api/model/web"
	"github.com/jejejery/go-restful-api/service/mocks"
	"github.com/gofiber/fiber/v2"
	gomock "go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupEmployeeTestApp(mockService *mocks.MockEmployeeService) *fiber.App {
	app := fiber.New()
	employeeController := NewEmployeeController(mockService)

	api := app.Group("/api")
	employees := api.Group("/employees")
	employees.Post("/", employeeController.Create)
	employees.Put("/:employeeId", employeeController.Update)
	employees.Delete("/:employeeId", employeeController.Delete)
	employees.Get("/:employeeId", employeeController.FindById)
	employees.Get("/", employeeController.FindAll)

	return app
}

func TestEmployeeController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockEmployeeService(ctrl)
	app := setupEmployeeTestApp(mockService)

	tests := []struct {
		name           string
		method         string
		url            string
		body           interface{}
		setupMock      func()
		expectedStatus int
		expectedBody   web.WebResponse
	}{
		{
			name:   "Update employee - success",
			method: "PUT",
			url:    "/api/employees/1",
			body:   web.EmployeeUpdateRequest{EmployeeID: "E001", Name: "Sayang"},
			setupMock: func() {
				mockService.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(web.EmployeeResponse{EmployeeID: "E001", Name: "Sayang"}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: web.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   web.EmployeeResponse{EmployeeID: "E001", Name: "Sayang"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			var reqBody []byte
			if tt.body != nil {
				reqBody, _ = json.Marshal(tt.body)
			}

			req := httptest.NewRequest(tt.method, tt.url, bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			var respBody web.WebResponse
			json.NewDecoder(resp.Body).Decode(&respBody)

			if dataMap, ok := respBody.Data.(map[string]interface{}); ok {
				respBody.Data = web.EmployeeResponse{
					EmployeeID:   dataMap["employee_id"].(string),
					Name: dataMap["name"].(string),
				}
			}

			assert.Equal(t, tt.expectedBody, respBody)
		})
	}
}