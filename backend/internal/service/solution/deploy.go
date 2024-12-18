package solution

import (
	"context"
	"fmt"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"go.uber.org/zap"
)

func (s serv) Deploy(ctx context.Context, deploy *model.Deploy) (int64, error) {
	solution, err := s.solutionRepository.Get(ctx, deploy.DeployInfo.SolutionId)
	if err != nil {
		return 0, err
	}
	logger.Info(
		"Deploy solution",
		zap.Int64("solution_id", deploy.DeployInfo.SolutionId),
		zap.String("solution", fmt.Sprintf("%+v", solution)))

	idPlaybook, _ := s.runnerService.RunPlaybook(ctx, *deploy, *solution)
	deploy.DeployInfo.IDPlaybook = idPlaybook

	id, err := s.deployRepository.Deploy(ctx, deploy)
	if err != nil {
		return 0, err
	}

	return id, nil
}
