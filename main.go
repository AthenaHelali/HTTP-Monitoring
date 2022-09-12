package main

import (
	"os"
	"log"

	"github.com/AthenaHelali/HTTP-Monitoring/internal/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := echo.New()
	routes.Register(app.Group("/users"))
	if err := app.Start(":"+port); err != nil {
		log.Println(err)
	}
}
