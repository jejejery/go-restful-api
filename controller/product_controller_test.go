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

func setupProductTestApp(mockService *mocks.MockProductService) *fiber.App {
	app := fiber.New()
	productController := NewProductController(mockService)

	api := app.Group("/api")
	products := api.Group("/products")
	products.Post("/", productController.Create)
	products.Put("/:productId", productController.Update)
	products.Delete("/:productId", productController.Delete)
	products.Get("/:productId", productController.FindById)
	products.Get("/", productController.FindAll)

	return app
}

func TestProductController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockProductService(ctrl)
	app := setupProductTestApp(mockService)

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
			name:   "Update product - success",
			method: "PUT",
			url:    "/api/products/1",
			body:   web.ProductUpdateRequest{Id: 0, Name: "Sayang"},
			setupMock: func() {
				mockService.EXPECT().
					Update(gomock.Any(), gomock.Any()).
					Return(web.ProductResponse{ProductID: 0, Name: "Sayang"}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: web.WebResponse{
				Code:   http.StatusOK,
				Status: "OK",
				Data:   web.ProductResponse{ProductID: 0, Name: "Sayang"},
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
				respBody.Data = web.ProductResponse{
					ProductID:   int(dataMap["product_id"].(float64)),
					Name: dataMap["name"].(string),
				}
			}

			assert.Equal(t, tt.expectedBody, respBody)
		})
	}
}