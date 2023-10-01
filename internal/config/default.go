package config

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
	"time"
)

func Default() Config {
	return Config{
		Debug: true,
		Database: Repository.Config{
			URL:               "mongodb://127.0.0.1:27017",
			Name:              "users",
			ConnectionTimeout: time.Second * 5,
		},
	}
}
