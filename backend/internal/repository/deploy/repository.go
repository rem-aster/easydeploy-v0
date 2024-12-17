package deploy

import (
	"github.com/s0vunia/platform_common/pkg/db"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository"
)

const (
	tableName = "deploy"

	idColumn         = "id"
	nameColumn       = "name"
	solutionIdColumn = "solution_id"
	sshAddressColumn = "ssh_address"
	sshKeyColumn     = "ssh_key"
	extraColumn      = "extra"
	idPlaybookColumn = "id_playbook"
	createdAtColumn  = "created_at"
	updatedAtColumn  = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new user repository.
func NewRepository(db db.Client) repository.DeployRepository {
	return &repo{db: db}
}
