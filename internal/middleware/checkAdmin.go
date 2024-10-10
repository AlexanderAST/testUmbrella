package middleware

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const userRole = "admin"

func CheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		headers := req.Header
		userRole := headers.Get("User-Role")

		if strings.ToLower(userRole) == "admin" {
			log.Print("red button user detected")
		}

		return next(c)
	}
}
