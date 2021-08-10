package main

import (
	"assignment/pkg/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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
}
