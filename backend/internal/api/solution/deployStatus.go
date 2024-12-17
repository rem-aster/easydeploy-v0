package solution

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
	"go.uber.org/zap"
)

func (i *Implementation) GetDeployStatus(ctx context.Context, req *desc.GetDeployStatusRequest) (*desc.GetDeployStatusResponse, error) {
	logger.Info("Get deploy status...")

	status, errS, err := i.solutionService.DeployStatus(ctx, req.GetId())
	if err != nil {
		logger.Error("failed to get deploy status", zap.Error(err))
		return nil, InternalError
	}
	return &desc.GetDeployStatusResponse{
		Status: status,
		Error:  errS,
	}, nil
}
