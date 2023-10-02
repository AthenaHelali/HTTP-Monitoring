package userhandler

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/auth"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) allUsers(c echo.Context) error {
	var req param.GetUsersRequest
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.Claims)
	req.UserID = claims.UserID

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	resp, err := h.UserSvc.GetAll(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cant get users")
	}

	return c.JSON(http.StatusCreated, resp)
}
