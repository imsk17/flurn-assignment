package main

import (
	"assignment/api/handlers"
	"assignment/pkg/entities"
	"assignment/pkg/loan"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN: "postgres://postgres:postgres@localhost/assignment",
		},
	))
	if err != nil {
		log.Panicf("Failed to connect to the database: %e", err)
	}
	if db.AutoMigrate(&entities.LoanRequest{}) != nil {
		log.Panicf("Failed to migrate the database: %e", err)
	}
	loanRepository := loan.NewRepository(db)
	handlers.SetupLoanRoutes(app, loanRepository)
	log.Panic(app.Listen(":8080"))
}
