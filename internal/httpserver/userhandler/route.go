package userhandler

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/register", h.userRegister)

	userGroup.POST("/login", h.userLogin)

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.Claims)
		},
		SigningKey: []byte("secret"),
	}
	userGroup.Use(echojwt.WithConfig(config))

	userGroup.GET("/all", h.allUsers)

}
