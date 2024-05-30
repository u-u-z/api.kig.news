package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	// import Price from ../models.go
)

func SubmitPrice(c echo.Context) error {

	var price Price

	if err := c.Bind(&price); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request payload",
		})
	}

	if err := c.Validate(&price); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	now := time.Now().UnixMilli()

	price.Time = now

	if err := db.Create(&price).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to save price",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"ok":      true,
		"price":   price.Price,
		"time":    now,
		"details": price.Details,
	})
}
