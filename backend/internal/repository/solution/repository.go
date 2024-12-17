package solution

import (
	"github.com/s0vunia/platform_common/pkg/db"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository"
)

const (
	tableName = "solution"

	idColumn          = "id"
	nameColumn        = "name"
	descriptionColumn = "description"
	createdAtColumn   = "created_at"
	updatedAtColumn   = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new user repository.
func NewRepository(db db.Client) repository.SolutionRepository {
	return &repo{db: db}
}
