package config

import (
	"github.com/AthenaHelali/HTTP-Monitoring/internal/Repository"
)

type Config struct {
	Debug    bool              `koanf:"debug"`
	Database Repository.Config `koanf:"database"`
}
