package deploy

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/platform_common/pkg/db"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/deploy/converter"
)

func (r *repo) Deploy(ctx context.Context, model *model.Deploy) (int64, error) {
	modelRepo := converter.ToDeployRepoFromModel(model)

	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, solutionIdColumn, idPlaybookColumn, sshAddressColumn, sshKeyColumn, extraColumn, createdAtColumn, updatedAtColumn).
		Values(modelRepo.DeployInfo.Name, modelRepo.DeployInfo.SolutionId, modelRepo.DeployInfo.IDPlaybook, modelRepo.DeployInfo.SSHAddress, modelRepo.DeployInfo.SSHKey, modelRepo.DeployInfo.Extra, modelRepo.CreatedAt, modelRepo.UpdatedAt).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "deploy_repository.Deploy",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
