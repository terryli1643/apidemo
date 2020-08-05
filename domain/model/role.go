package model

import (
	"time"
)

type Role struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time

	RoleCode    string
	RoleName    string
	Description string
	State       AccountState
}
