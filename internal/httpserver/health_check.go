package httpserver

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a App) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "everything is good",
	})

}
