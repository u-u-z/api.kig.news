package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func KeyAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("Authorization")

			if key == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Key")
			}

			var auth Auth
			db.First(&auth, "key = ?", key)

			if auth.ID == 0 {
				return echo.NewHTTPError(http.StatusForbidden, "Invalid Key")
			}

			if auth.RemainingPoints <= 0 {
				return echo.NewHTTPError(http.StatusForbidden, "You do not have permission to access this resource")
			}

			c.Set("auth", auth)

			return next(c)
		}
	}
}
