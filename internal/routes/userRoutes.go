package routes

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/controller"
	"github.com/labstack/echo/v4"
)

func Register(g *echo.Group) {
	g.POST("/signup", controller.SignUp)
	g.POST("/login", controller.Login)
	//g.Use(middleware.Authentication())
	g.GET("/history", controller.GetHistory)
	g.GET("/alerts", controller.GetAlerts)

}
