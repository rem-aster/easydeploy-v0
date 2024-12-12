package solution

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/platform_common/pkg/db"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/solution/converter"
	modelRepo "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/solution/model"
)

func (r *repo) List(ctx context.Context) ([]*model.Solution, error) {
	builderSelectOne := sq.Select(idColumn, nameColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "solution_repository.List",
		QueryRaw: query,
	}
	var solutionsRepo []*modelRepo.Solution
	err = r.db.DB().ScanAllContext(ctx, &solutionsRepo, q, args...)
	if err != nil {
		return nil, err
	}

	solutions := make([]*model.Solution, 0, len(solutionsRepo))
	for _, solution := range solutionsRepo {
		solutions = append(solutions, converter.ToSolutionFromRepo(solution))
	}
	return solutions, nil
}
