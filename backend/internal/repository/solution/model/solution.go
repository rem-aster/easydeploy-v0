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
	Name        string `db:"name"`
	Description string `db:"description"`
	Playbook    string `db:"playbook"`
}

// DeployStatus represents a status of deployment
