package dao

import "github.com/terryli1643/apidemo/domain/model"

func (dao *AdminDao) FindAdmin(condition model.SearchAdminCondition, rowbound model.RowBound) (result []model.Admin, count int64, err error) {
	var admin model.Admin
	var role model.Role

	if condition.Account != "" {
		admin.Account = condition.Account
	}
	if condition.RoleID != 0 {
		role.ID = condition.RoleID
	}
	if condition.Department != "" {
		admin.Department = condition.Department
	}
	if condition.State != 0 {
		admin.State = condition.State
	}

	chain := dao.db.Model(&role).Preload("Roles")
	if condition.RoleID != 0 {
		chain = chain.Joins("left join `role_user_rel` on `admin`.id = `role_user_rel`.user_id").Where("`role_user_rel`.role_id = ?", condition.RoleID)
	}
	chain = chain.Model(&admin).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset)
	err = chain.Find(&result, admin).Error
	return
}
