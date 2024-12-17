package env

import (
	"errors"
	"net"
	"os"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/config"
)

const (
	authHostEnvName = "RUNNER_HOST"
	authPortEnvName = "RUNNER_PORT"
)

type runnerConfig struct {
	host string
	port string
}

// NewRunnerServiceConfig initializes a gRPC configuration.
func NewRunnerServiceConfig() (config.RunnerServiceConfig, error) {
	host := os.Getenv(authHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(authPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &runnerConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *runnerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
