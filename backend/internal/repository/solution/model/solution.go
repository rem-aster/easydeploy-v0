package model

import (
	"database/sql"
	"time"
)

// Solution represents a user entity with ID, Info, CreatedAt, and UpdatedAt fields.
type Solution struct {
	ID        int64        `db:"id"`
	Info      SolutionInfo `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// SolutionInfo represents a user info entity with Status fields.
type SolutionInfo struct {
	Status DeployStatus `db:"status"`
}

// DeployStatus represents a status of deployment
type DeployStatus int32

const (
	// DeployUnknown represents an unknown deploy status
	DeployUnknown DeployStatus = 0
)
