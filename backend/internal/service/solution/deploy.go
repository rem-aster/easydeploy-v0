package solution

import "context"

func (s serv) Deploy(ctx context.Context) error {
	return s.solutionRepository.Deploy(ctx)
}
