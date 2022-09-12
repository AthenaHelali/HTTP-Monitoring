package controller

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/request"
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	var req = request.user{

	}
}
func Login(c echo.Context) error {
	return nil

}
func GetHistory(c echo.Context) error {
	return echo.ErrBadRequest

}
func GetAlerts(c echo.Context) error {
	return nil
}
