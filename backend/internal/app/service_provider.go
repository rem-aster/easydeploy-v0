package app

import (
	"context"

	"github.com/s0vunia/platform_common/pkg/closer"
	"github.com/s0vunia/platform_common/pkg/db"
	"github.com/s0vunia/platform_common/pkg/db/pg"
	"github.com/s0vunia/platform_common/pkg/db/transaction"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/api/solution"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	"go.uber.org/zap"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/config/env"
	deployRepository "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/deploy"
	solutionRepository "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/solution"
	solutionService "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service/solution"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/config"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service"

	"github.com/avast/retry-go"
)

type serviceProvider struct {
	pgConfig         config.PGConfig
	grpcConfig       config.GRPCConfig
	httpConfig       config.HTTPConfig
	prometheusConfig config.PrometheusConfig
	loggerConfig     config.LoggerConfig

	dbClient  db.Client
	txManager db.TxManager

	solutionRepository repository.SolutionRepository
	deployRepository   repository.DeployRepository

	solutionService service.SolutionService

	solutionImpl *solution.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			logger.Fatal(
				"failed to get pg config",
				zap.Error(err),
			)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			logger.Fatal(
				"failed to get grpc config",
				zap.Error(err),
			)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			logger.Fatal(
				"failed to get http config",
				zap.Error(err),
			)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) PrometheusConfig() config.PrometheusConfig {
	if s.prometheusConfig == nil {
		cfg, err := env.NewPrometheusConfig()
		if err != nil {
			logger.Fatal(
				"failed to get prometheus config",
				zap.Error(err),
			)
		}

		s.prometheusConfig = cfg
	}

	return s.prometheusConfig
}

func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := env.NewLoggerConfig()
		if err != nil {
			logger.Fatal(
				"failed to get logger config",
				zap.Error(err),
			)
		}
		s.loggerConfig = cfg
	}
	return s.loggerConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		var cl db.Client
		err := retry.Do(
			func() error {
				var err error
				cl, err = pg.New(ctx, s.PGConfig().DSN())
				return err
			},
			retry.Attempts(3),
			retry.OnRetry(func(n uint, err error) {
				logger.Info("Retrying request after error: %v",
					zap.Error(err))
			}),
		)
		if err != nil {
			logger.Fatal(
				"failed to get db client",
				zap.Error(err),
			)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			logger.Fatal(
				"failed to ping db",
				zap.Error(err),
			)
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) SolutionRepository(ctx context.Context) repository.SolutionRepository {
	if s.solutionRepository == nil {
		s.solutionRepository = solutionRepository.NewRepository(s.DBClient(ctx))
	}

	return s.solutionRepository
}

func (s *serviceProvider) DeployRepository(ctx context.Context) repository.DeployRepository {
	if s.deployRepository == nil {
		s.deployRepository = deployRepository.NewRepository(s.DBClient(ctx))
	}

	return s.deployRepository
}

func (s *serviceProvider) SolutionService(ctx context.Context) service.SolutionService {
	if s.solutionService == nil {
		s.solutionService = solutionService.NewService(
			s.SolutionRepository(ctx),
			s.DeployRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.solutionService
}

func (s *serviceProvider) SolutionImpl(ctx context.Context) *solution.Implementation {
	if s.solutionImpl == nil {
		s.solutionImpl = solution.NewImplementation(s.SolutionService(ctx))
	}

	return s.solutionImpl
}
