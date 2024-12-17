package repository

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
)

// SolutionRepository represents a repository for solution entities.
type SolutionRepository interface {
	List(ctx context.Context) ([]*model.Solution, error)
}

// DeployRepository represents a repository for deploy entities.
type DeployRepository interface {
	Deploy(ctx context.Context, deploy *model.Deploy) (int64, error)
	DeployStatus(ctx context.Context) error
}
