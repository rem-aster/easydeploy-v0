package runnerService

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
)

// RunnerService interface
type RunnerService interface {
	RunPlaybook(ctx context.Context, deploy model.Deploy, solution model.Solution) (string, string)
	GetTaskStatus(ctx context.Context, id string) (string, string)
}
