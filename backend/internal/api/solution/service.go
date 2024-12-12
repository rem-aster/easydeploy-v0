package solution

import (
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
)

// Implementation represents a deploy API implementation.
type Implementation struct {
	desc.UnimplementedSolutionV1Server
	solutionService service.SolutionService
}

// NewImplementation creates a new solution API implementation.
func NewImplementation(solutionService service.SolutionService) *Implementation {
	return &Implementation{
		solutionService: solutionService,
	}
}
