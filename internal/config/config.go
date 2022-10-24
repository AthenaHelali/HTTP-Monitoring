package config

import "github.com/AthenaHelali/HTTP-Monitoring/internal/db"

type Config struct {
	Debug    bool      `koanf:"debug"`
	Database db.Config `koanf:"database"`
}
