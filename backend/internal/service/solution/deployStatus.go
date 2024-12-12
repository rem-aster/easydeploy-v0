package solution

import "context"

func (s serv) DeployStatus(ctx context.Context) error {
	return s.solutionRepository.DeployStatus(ctx)
}
