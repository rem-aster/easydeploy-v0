package converter

import (
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	modelRepo "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/deploy/model"
)

func ToDeployRepoFromModel(deploy *model.Deploy) *modelRepo.Deploy {
	return &modelRepo.Deploy{
		ID:         deploy.ID,
		DeployInfo: ToDeployInfoRepoFromModel(&deploy.DeployInfo),
		CreatedAt:  deploy.CreatedAt,
		UpdatedAt:  deploy.UpdatedAt,
	}
}

func ToDeployInfoRepoFromModel(deployInfo *model.DeployInfo) modelRepo.DeployInfo {
	return modelRepo.DeployInfo{
		SolutionId: deployInfo.SolutionId,
		Status:     modelRepo.DeployStatus(deployInfo.Status),
		Name:       deployInfo.Name,
		SSHAddress: deployInfo.SSHAddress,
		SSHKey:     deployInfo.SSHKey,
		Extra:      deployInfo.Extra,
	}
}
