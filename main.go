package main

import (
	"log"
	"os"

	"github.com/AthenaHelali/HTTP-Monitoring/internal/config"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/db"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/handler"
	"github.com/AthenaHelali/HTTP-Monitoring/internal/store"
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

	db, err := db.New(cfg.Database)
	if err != nil {
		logger.Named("db").Fatal("cannot create a db instance", zap.Error(err))
	}
	userStore := store.NewUserMongoDB(
		db, logger.Named("user"),
	)
	h := handler.App{
		Store:  *userStore,
		Logger: logger.Named("user"),
	}
	h.Register(app)
	h.Start()
	if err := app.Start(":" + port); err != nil {
		log.Println(err)
	}
}
