package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (s *server) getTime(c echo.Context) error {

	t1 := time.Now()
	t2 := time.Date(2025, time.January, 01, 0, 0, 0, 0, time.UTC)

	duration := t2.Sub(t1)

	result := int(duration.Hours() / 24)

	return c.JSON(http.StatusOK, map[string]interface{}{"days from 1 jan 2025": result})

}
