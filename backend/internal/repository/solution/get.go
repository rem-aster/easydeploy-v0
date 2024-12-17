package solution

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/platform_common/pkg/db"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/solution/converter"
	modelRepo "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/solution/model"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.Solution, error) {
	builderSelectOne := sq.Select(nameColumn, playbookColumn, descriptionColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "solution_repository.Solution",
		QueryRaw: query,
	}

	var res modelRepo.Solution

	err = r.db.DB().ScanOneContext(ctx, &res, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToSolutionFromRepo(&res), nil
}
