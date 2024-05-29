package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	e  *echo.Echo
)

func main() {
	DBInit()
	HTTPInit()
	RouteInit()

	e.Logger.Fatal(e.Start(":8080"))
}
