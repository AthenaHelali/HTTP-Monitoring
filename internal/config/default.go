package config

import (
	"time"

	"github.com/AthenaHelali/HTTP-Monitoring/internal/db"
)

func Default() Config {
	return Config{
		Debug: true,
		Database: db.Config{
			URL:               "mongodb://127.0.0.1:27017",
			Name:              "users",
			ConnectionTimeout: time.Second * 5,
		},
	}
}
