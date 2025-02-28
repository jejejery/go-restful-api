package main

import (
	"github.com/jejejery/go-restful-api/app"
	"github.com/jejejery/go-restful-api/controller"
	"github.com/jejejery/go-restful-api/helper"
	"github.com/jejejery/go-restful-api/model/domain"
	"github.com/jejejery/go-restful-api/repository"
	"github.com/jejejery/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	server := fiber.New()

	// Initialize Database
	db := app.NewDB()

	// Run Auto Migration (Opsional, bisa dihapus jika tidak diperlukan)
	err := db.AutoMigrate(&domain.Category{})
	helper.PanicIfError(err)

	// Initialize Validator
	validate := validator.New()

	// Initialize Repository, Service, and Controller
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// Setup Routes
	app.NewRouter(server, categoryController)

	// Start Server
	log.Println("Server running on port 8081")
	err = server.Listen(":8081")
	helper.PanicIfError(err)
}
