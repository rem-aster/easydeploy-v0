package converter

import (
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToSolutionFromService converts a object from the model package to a Solution object from the desc package.
func ToSolutionFromService(solution *model.Solution) *desc.Solution {
	var updatedAt *timestamppb.Timestamp
	if solution.UpdatedAt.Valid {
		updatedAt = timestamppb.New(solution.UpdatedAt.Time)
	}
	return &desc.Solution{
		Id: solution.ID,
		Info: &desc.SolutionInfo{
			Name:        solution.Info.Name,
			Description: solution.Info.Description,
		},
		CreatedAt: timestamppb.New(solution.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
