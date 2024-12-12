package solution

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
)

func (i *Implementation) DeployStatus(ctx context.Context, req *desc.GetDeployStatusRequest) (*desc.GetDeployStatusResponse, error) {
	logger.Info("Get deploy status...")

	// TODO
	err := i.solutionService.DeployStatus(ctx)
	return &desc.GetDeployStatusResponse{}, err
}
