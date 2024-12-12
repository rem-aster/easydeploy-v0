package solution

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/converter"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
)

// List Solutions.
func (i *Implementation) List(ctx context.Context, req *desc.ListRequest) (*desc.ListResponse, error) {
	logger.Info("List solutions...")
	solutions, err := i.solutionService.List(ctx)
	if err != nil {
		return nil, err
	}

	var descSolutions []*desc.Solution
	for _, solution := range solutions {
		descSolutions = append(descSolutions, converter.ToSolutionFromService(solution))
	}
	return &desc.ListResponse{
		Solutions: descSolutions,
	}, nil
}
