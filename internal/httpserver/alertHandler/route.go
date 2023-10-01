package alertHandler

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUrlRoutes(e *echo.Echo) {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.Claims)
		},
		SigningKey: []byte("secret"),
	}

	urlGroup := e.Group("/alerts")
	urlGroup.Use(echojwt.WithConfig(config))

	urlGroup.GET("/get", h.GetAlert)
}
