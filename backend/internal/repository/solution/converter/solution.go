package converter

import (
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	modelRepo "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/solution/model"
)

// ToSolutionFromRepo converts a Solution object from the model package to a Solution object from the repository package.
func ToSolutionFromRepo(solution *modelRepo.Solution) *model.Solution {
	return &model.Solution{
		ID:        solution.ID,
		Info:      ToSolutionInfoFromRepo(solution.Info),
		CreatedAt: solution.CreatedAt,
		UpdatedAt: solution.UpdatedAt,
	}
}

// ToSolutionInfoFromRepo converts a SolutionInfo object from the model package to a SolutionInfo object from the repository package.
func ToSolutionInfoFromRepo(info modelRepo.SolutionInfo) model.SolutionInfo {
	return model.SolutionInfo{
		Name:        info.Name,
		Description: info.Description,
		Playbook:    info.Playbook,
	}
}
