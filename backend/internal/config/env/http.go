package env

import (
	"errors"
	"net"
	"os"
	"strconv"
	"time"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/config"
)

const (
	httpHostEnvName          = "HTTP_HOST"
	httpPortEnvName          = "HTTP_PORT"
	readHeaderTimeoutEnvName = "READ_HEADER_TIMEOUT_SEC"
)

type httpConfig struct {
	host              string
	port              string
	readHeaderTimeout time.Duration
}

// NewHTTPConfig - creates new http config
func NewHTTPConfig() (config.HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	timeout, err := strconv.Atoi(os.Getenv(readHeaderTimeoutEnvName))
	if len(port) == 0 && err != nil {
		return nil, errors.New("read header timeout not found")
	}
	readHeaderTimeout := time.Second * time.Duration(timeout)

	return &httpConfig{
		host:              host,
		port:              port,
		readHeaderTimeout: readHeaderTimeout,
	}, nil
}

// Address - creates address from host and port
func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

// ReadHeaderTimeout - returns read header timeout
func (cfg *httpConfig) ReadHeaderTimeout() time.Duration {
	return cfg.readHeaderTimeout
}
