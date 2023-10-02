package main

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository/mongo"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/httpserver"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/httpserver/alertHandler"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/httpserver/urlhandler"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/httpserver/userhandler"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/service/alert"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/service/monitor"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/service/url"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/service/user"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/validator/uservalidator"
	"log"
	"os"

	"github.com/AthenaHelali/HTTP-Monitoring/internal/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}
	cfg := config.New()
	app := echo.New()

	var logger *zap.Logger
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatal(err)
	}

	db, err := Repository.New(cfg.Database)
	if err != nil {
		logger.Named("db").Fatal("cannot create a db instance", zap.Error(err))
	}
	userStore := mongo.NewUserMongoDB(
		db, logger.Named("user"),
	)
	userHandler := userhandler.Handler{
		UserSvc:       user.New(userStore),
		UserValidator: uservalidator.New(userStore),
	}
	urlHandler := urlhandler.Handler{
		UrlSvc: url.New(userStore),
	}
	alertHandler := alertHandler.Handler{
		UrlSvc: alert.New(userStore),
	}
	monitoring := monitor.New(userStore)
	h := httpserver.App{
		Store:        *userStore,
		Logger:       logger.Named("user"),
		UserHandler:  userHandler,
		UrlHandler:   urlHandler,
		AlertHandler: alertHandler,
		Monitoring:   monitoring,
	}
	h.Serve(app)
	h.Start()
	if err := app.Start(":" + port); err != nil {
		log.Println(err)
	}
}
