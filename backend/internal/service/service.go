package service

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
)

// SolutionService represents a service for solution entities.
type SolutionService interface {
	List(ctx context.Context) ([]*model.Solution, error)
	Deploy(ctx context.Context) error
	DeployStatus(ctx context.Context) error
}
