package solution

import (
	"context"
)

func (s serv) DeployStatus(ctx context.Context, id int64) (string, string, error) {
	deploy, err := s.deployRepository.Get(ctx, id)
	if err != nil {
		return "", "", err
	}
	status, errS := s.runnerService.GetTaskStatus(ctx, deploy.DeployInfo.IDPlaybook)
	return status, errS, nil
}
