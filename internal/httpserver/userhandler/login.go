package userhandler

import (
	"errors"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/auth"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h Handler) userLogin(c echo.Context) error {
	var req param.LoginRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	if err := h.UserValidator.ValidateLoginRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	user, err := h.UserSvc.Login(req)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	claims := &auth.Claims{
		UserID:           user.ID,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
