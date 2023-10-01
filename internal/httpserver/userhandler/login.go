package userhandler

import (
	"errors"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/param"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h Handler) userLogin(c echo.Context) error {
	var req param.LoginRequest

	if err := c.Bind(&req); err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, "cant bind request")
	}

	user, err := h.UserSvc.Login(req)
	if err != nil {
		var ErrNotFound Repository.UserNotFoundError
		if ok := errors.As(err, &ErrNotFound); ok {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// Set custom claims
	claims := &JWTCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

type JWTCustomClaims struct {
	ID string `json:"name"`
	jwt.StandardClaims
}
