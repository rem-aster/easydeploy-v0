package converter

import (
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
)

func ToDeployFromDesc(deploy *desc.DeployRequest) *model.Deploy {
	return &model.Deploy{
		DeployInfo: model.DeployInfo{
			SolutionId: deploy.GetSolutionId(),
			SSHAddress: deploy.GetSshAddress(),
			SSHKey:     deploy.GetSshKey(),
			Extra:      deploy.GetExtraVars(),
		},
	}
}
