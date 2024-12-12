package solution

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
)

func (s serv) List(ctx context.Context) ([]*model.Solution, error) {
	return s.solutionRepository.List(ctx)
}
