package solution

import (
	"github.com/s0vunia/platform_common/pkg/db"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/client/runnerService"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service"
)

type serv struct {
	solutionRepository repository.SolutionRepository
	deployRepository   repository.DeployRepository
	runnerService      runnerService.RunnerService
	txManager          db.TxManager
}

// NewService creates a new user service.
func NewService(
	solutionRepository repository.SolutionRepository,
	deployRepository repository.DeployRepository,
	runnerService runnerService.RunnerService,
	txManager db.TxManager,
) service.SolutionService {
	return &serv{
		solutionRepository: solutionRepository,
		deployRepository:   deployRepository,
		runnerService:      runnerService,
		txManager:          txManager,
	}
}
