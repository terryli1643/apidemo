package model

import "time"

type RoleUserRel struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time

	UserID int64
	RoleID int64
}
