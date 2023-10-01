package userhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) allUsers(c echo.Context) error {
	resp, err := h.UserSvc.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cant get users")
	}

	return c.JSON(http.StatusCreated, resp)
}
