package model

import (
	"time"
)

type Admin struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	//用户类型
	UserType UserType
	//帐号
	Account string
	//工号
	JobNumber string
	//登录密码
	LoginPassword string
	//安全密码
	SecurePassword string
	//部门
	Department string
	//创建人ID
	Creator int64
	//创建人账户
	CreatorAccount string
	//备注
	Remark string
	//帐号状态,  正常/锁定
	State AccountState
	//最后登录时间
	LastLoginAt *time.Time
	//最后登录IP
	LastLoginIP string
	//最后登录ip地址
	LastLoginIPAddr string
	//角色
	Roles []*Role `gorm:"many2many:role_user_rel;association_jointable_foreignkey:role_id;jointable_foreignkey:user_id;"`
	//IM 账号
	IMAccount string
}

type SearchAdminCondition struct {
	Account    string
	RoleID     int64
	Department string
	State      AccountState
}

func (admin *Admin) GetUsername() string {
	return admin.Account
}
func (admin *Admin) GetPassword() string {
	return admin.LoginPassword
}
func (admin *Admin) IsAccountExpired() bool {
	return false
}
func (admin *Admin) IsAccountLocked() bool {
	return admin.State == AccountDisable
}
func (admin *Admin) IsCredentialsExpired() bool {
	return false
}
