package solution

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
)

func (i *Implementation) Deploy(ctx context.Context, req *desc.DeployRequest) (*desc.DeployResponse, error) {
	logger.Info("Deploy solution...")

	// TODO
	err := i.solutionService.Deploy(ctx)
	return &desc.DeployResponse{}, err
}
