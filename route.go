package main

import (
	"github.com/labstack/echo/v4"
)

func RouteInit() {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.POST("/price", SubmitPrice)
}
