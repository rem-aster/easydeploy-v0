package repository

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
)

// SolutionRepository represents a repository for solution entities.
type SolutionRepository interface {
	List(ctx context.Context) ([]*model.Solution, error)
	Deploy(ctx context.Context) error
	DeployStatus(ctx context.Context) error
}
