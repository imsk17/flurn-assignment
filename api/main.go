package main

import (
	"assignment/api/handlers"
	"assignment/pkg/entities"
	"assignment/pkg/loan"
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm/logger"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PORT = "8080"

func init() {
	port := os.Getenv("PORT")
	if port != "" {
		PORT = port
	}
}

func main() {
	app := fiber.New()
	app.Use(fiberLogger.New())
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN: "postgres://postgres:postgres@localhost/assignment",
		},
	), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Panicf("Failed to connect to the database: %e", err)
	}
	// Our ORM right now does not support auto migration
	// for enum types and hence we have to migrate it ourselves.
	//! IMPORTANT: Should be run only once
	_ = db.Exec("CREATE TYPE status AS ENUM ('New', 'Approved', 'Rejected', 'Cancelled')").Error

	if db.AutoMigrate(&entities.LoanRequest{}) != nil {
		log.Panicf("Failed to migrate the database: %e", err)
	}
	loanRepository := loan.NewRepository(db)
	handlers.SetupLoanRoutes(app, loanRepository)
	log.Panic(app.Listen(":" + PORT))
}
