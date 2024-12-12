package model

import (
	"database/sql"
	"time"
)

// Solution represents a user entity with ID, Info, CreatedAt, and UpdatedAt fields.
type Solution struct {
	ID        int64
	Info      SolutionInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// SolutionInfo represents a user info entity with Status fields.
type SolutionInfo struct {
	Name   string
	Status DeployStatus
}

// DeployStatus represents a status of deployment
type DeployStatus int32

const (
	// DeployUnknown represents an unknown deploy status
	DeployUnknown DeployStatus = 0
)
