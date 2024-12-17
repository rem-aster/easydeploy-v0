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
		IDPlaybook: deployInfo.IDPlaybook,
	}
}

func ToDeployModelFromRepo(deploy *modelRepo.Deploy) *model.Deploy {
	return &model.Deploy{
		ID:         deploy.ID,
		DeployInfo: ToDeployInfoModelFromRepo(&deploy.DeployInfo),
		CreatedAt:  deploy.CreatedAt,
		UpdatedAt:  deploy.UpdatedAt,
	}
}

func ToDeployInfoModelFromRepo(deployInfo *modelRepo.DeployInfo) model.DeployInfo {
	return model.DeployInfo{
		SolutionId: deployInfo.SolutionId,
		Status:     model.DeployStatus(deployInfo.Status),
		Name:       deployInfo.Name,
		SSHAddress: deployInfo.SSHAddress,
		SSHKey:     deployInfo.SSHKey,
		Extra:      deployInfo.Extra,
		IDPlaybook: deployInfo.IDPlaybook,
	}
}
