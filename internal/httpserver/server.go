package httpserver

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository/mongo"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/httpserver/alertHandler"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/httpserver/urlhandler"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/httpserver/userhandler"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/service/monitor"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type App struct {
	Store        mongo.UserMongodb
	Logger       *zap.Logger
	UserHandler  userhandler.Handler
	UrlHandler   urlhandler.Handler
	AlertHandler alertHandler.Handler
	Monitoring   monitor.Service
}

func (a App) Serve(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health-check", a.healthCheck)

	a.UserHandler.SetUserRoutes(e)
	a.UrlHandler.SetUrlRoutes(e)
	a.AlertHandler.SetUrlRoutes(e)
}

func (a App) Start() {
	go func() {
		a.Monitoring.MonitorAllUsers()
	}()
}
