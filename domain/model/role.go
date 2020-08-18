package model

import (
	"time"
)

type Role struct {
	ID        int64      `json:"id,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	RoleCode    string       `json:"roleCode,omitempty"`
	RoleName    string       `json:"roleName,omitempty"`
	Description string       `json:"description,omitempty"`
	State       AccountState `json:"state,omitempty"`
}
