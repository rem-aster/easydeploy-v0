package solution

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
	"go.uber.org/zap"
)

func (i *Implementation) DeployStatus(ctx context.Context, req *desc.GetDeployStatusRequest) (*desc.GetDeployStatusResponse, error) {
	logger.Info("Get deploy status...")

	// TODO
	err := i.solutionService.DeployStatus(ctx)
	if err != nil {
		logger.Error("failed to deploy solution", zap.Error(err))
		return nil, InternalError
	}
	return &desc.GetDeployStatusResponse{}, nil
}
