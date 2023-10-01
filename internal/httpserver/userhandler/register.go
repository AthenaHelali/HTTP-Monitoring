package userhandler

import (
	"errors"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userRegister(c echo.Context) error {
	var req param.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	resp, err := h.UserSvc.Register(req)

	if err != nil {
		var errDuplicate Repository.DuplicateUserError
		if ok := errors.As(err, &errDuplicate); ok {
			return echo.NewHTTPError(http.StatusBadRequest, "user already exists")
		}

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, resp)
}
