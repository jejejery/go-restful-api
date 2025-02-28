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

func setupCustomerTestApp(mockService *mocks.MockCustomerService) *fiber.App {
	app := fiber.New()
	customerController := NewCustomerController(mockService)

	api := app.Group("/api")
	customers := api.Group("/customers")
	customers.Post("/", customerController.Create)
	customers.Put("/:customerId", customerController.Update)
	customers.Delete("/:customerId", customerController.Delete)
	customers.Get("/:customerId", customerController.FindById)
	customers.Get("/", customerController.FindAll)

	return app
}

func TestCustomerController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockCustomerService(ctrl)
	app := setupCustomerTestApp(mockService)

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
			name:   "Update customer - success",
			method: "PUT",
			url:    "/api/customers/1",
			body:   web.CustomerUpdateRequest{CustomerID: "C001", Name: "Sayang"},
			setupMock: func() {
				mockService.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(web.CustomerResponse{CustomerID: "C001", Name: "Sayang"}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: web.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   web.CustomerResponse{CustomerID: "C001", Name: "Sayang"},
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
				respBody.Data = web.CustomerResponse{
					CustomerID:   dataMap["customer_id"].(string),
					Name: dataMap["name"].(string),
				}
			}

			assert.Equal(t, tt.expectedBody, respBody)
		})
	}
}