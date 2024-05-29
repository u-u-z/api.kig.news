package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		// break the program if DATABASE_URL is not set
		panic("DATABASE_URL is not set")
	}

	db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	// Auto migrate the User model
	db.AutoMigrate(&Price{}, &Product{}, &Auth{})
}

func HTTPInit() {
	e = echo.New()
}

func RouteInit() {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
}
