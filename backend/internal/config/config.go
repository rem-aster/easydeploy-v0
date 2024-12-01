package config

import (
	"time"

	"github.com/joho/godotenv"
)

// GRPCConfig - config for GRPC server
type GRPCConfig interface {
	Address() string
}

// PGConfig - config for Postgres
type PGConfig interface {
	DSN() string
}

// HTTPConfig - config for HTTP server
type HTTPConfig interface {
	Address() string
	ReadHeaderTimeout() time.Duration
}

// SwaggerConfig - config for swagger
type SwaggerConfig interface {
	Address() string
}

// PrometheusConfig - config for prometheus
type PrometheusConfig interface {
	Address() string
}

// StorageConfig - config for storage
type StorageConfig interface {
	Mode() string
}

// LoggerConfig - config for logger
type LoggerConfig interface {
	Level() string
	FileName() string
	MaxSize() int
	MaxAge() int
	MaxBackups() int
}

// Load - loads config from .env
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
