package model

import (
	"database/sql"
	"time"
)

type Deploy struct {
	ID         string       `db:"id"`
	DeployInfo DeployInfo   `db:""`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
}

type DeployInfo struct {
	SolutionId int64        `db:"solution_id"`
	Status     DeployStatus `db:"status"`
	Name       string       `db:"name"`
	SSHAddress string       `db:"ssh_address"`
	SSHKey     []byte       `db:"ssh_key"`
	Extra      interface{}  `db:"extra"`
}
type DeployStatus int32

const (
	// DeployUnknown represents an unknown deploy status
	DeployUnknown DeployStatus = 0
)
