package model

import (
	"database/sql"
	"time"
)

type Deploy struct {
	ID         string
	DeployInfo DeployInfo
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}

type DeployInfo struct {
	SolutionId int64
	Status     DeployStatus
	Name       string
	SSHAddress string
	SSHKey     string
	Extra      interface{}
	IDPlaybook string
}
type DeployStatus int32

const (
	// DeployUnknown represents an unknown deploy status
	DeployUnknown DeployStatus = 0
)
