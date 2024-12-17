package deploy

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/platform_common/pkg/db"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/deploy/converter"
	modelRepo "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/deploy/model"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.Deploy, error) {
	builderSelectOne := sq.Select(nameColumn, solutionIdColumn, sshAddressColumn, sshKeyColumn, extraColumn, idPlaybookColumn, createdAtColumn, updatedAtColumn).
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

	var res modelRepo.Deploy

	err = r.db.DB().ScanOneContext(ctx, &res, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToDeployModelFromRepo(&res), nil
}
