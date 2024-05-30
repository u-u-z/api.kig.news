package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() {
	var err error

	if DATABASE_URL == "" {
		panic("ERROR: DB DATABASE_URL is not set")
	}

	db, err = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})

	if err != nil {
		panic("ERROR: Failed to connect to database")
	}

	db.AutoMigrate(&Price{}, &Product{}, &Auth{})
}

func HTTPServerInit() {
	e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}
}
