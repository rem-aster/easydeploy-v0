package solution

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
)

func (s serv) Deploy(ctx context.Context, deploy *model.Deploy) (int64, error) {
	id, err := s.deployRepository.Deploy(ctx, deploy)
	if err != nil {
		return 0, err
	}
	// TODO - runner grpc call

	return id, nil
}
