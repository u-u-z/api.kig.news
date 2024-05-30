package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	DATABASE_URL = os.Getenv("DATABASE_URL")
	SERVER_PORT  = os.Getenv("PORT")
)

var (
	db *gorm.DB
	e  *echo.Echo
)

func main() {

	DBInit()
	HTTPServerInit()
	RouteInit()

	e.Logger.Fatal(e.Start(":" + SERVER_PORT))
}
