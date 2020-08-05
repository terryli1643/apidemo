package model

import "time"

type RolePrivsRel struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time

	RoleID      int64
	PrivilegeID int64
}
