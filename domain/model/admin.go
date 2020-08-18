package model

import (
	"time"
)

type Admin struct {
	ID        int64      `json:"id,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	//用户类型
	UserType UserType `json:"userType,omitempty"`
	//帐号
	Account string `json:"account,omitempty"`
	//工号
	JobNumber string `json:"jobNumber,omitempty"`
	//登录密码
	LoginPassword string `json:"loginPassword,omitempty"`
	//安全密码
	SecurePassword string `json:"securePassword,omitempty"`
	//部门
	Department string `json:"department,omitempty"`
	//创建人ID
	Creator int64 `json:"creator,omitempty"`
	//创建人账户
	CreatorAccount string `json:"creatorAccount,omitempty"`
	//备注
	Remark string `json:"remark,omitempty"`
	//帐号状态,  正常/锁定
	State AccountState `json:"state,omitempty"`
	//最后登录时间
	LastLoginAt *time.Time `json:"lastLoginAt,omitempty"`
	//最后登录IP
	LastLoginIP string `json:"lastLoginIP,omitempty"`
	//最后登录ip地址
	LastLoginIPAddr string `json:"lastLoginIPAddr,omitempty"`
	//角色
	Roles []*Role `gorm:"many2many:role_user_rel;association_jointable_foreignkey:role_id;jointable_foreignkey:user_id;" json:"roles,omitempty"`
	//IM 账号
	IMAccount string `json:"imAccount,omitempty"`
}

type SearchAdminCondition struct {
	Account    string
	RoleID     int64
	Department string
	State      AccountState
}

func (admin Admin) GetID() int64 {
	return admin.ID
}

func (admin Admin) GetUsername() string {
	return admin.Account
}
func (admin Admin) GetPassword() string {
	return admin.LoginPassword
}
func (admin Admin) IsAccountLocked() bool {
	return admin.State == AccountDisable
}
