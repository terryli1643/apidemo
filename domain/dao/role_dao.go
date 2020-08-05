package dao

import (
	"errors"

	"github.com/terryli1643/apidemo/domain/model"
	"gorm.io/gorm"
)

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{
		db: db,
	}
}

func (dao *RoleDao) Create(m *model.Role) error {
	return dao.db.Create(m).Error
}

func (dao *RoleDao) Find(m *model.Role) (result []model.Role, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *RoleDao) FindOne(m *model.Role) error {
	return dao.db.First(m, m).Error
}

func (dao *RoleDao) FindLast(m *model.Role) error {
	return dao.db.Last(m, m).Error
}

func (dao *RoleDao) FindPage(m *model.Role, rowbound model.RowBound, desc bool) (result []model.Role, count int64, err error) {
	db := dao.db
	if desc {
		db = db.Order("id desc")
	}
	err = db.Model(m).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *RoleDao) Get(m *model.Role) error {
	if m.ID == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *RoleDao) BatchGet(idbatch []int64) (result []model.Role, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&model.Role{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *RoleDao) GetForUpdate(m *model.Role) error {
	if m.ID == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *RoleDao) Save(m *model.Role) error {
	return dao.db.Save(m).Error
}

func (dao *RoleDao) Delete(m *model.Role) error {
	if m.ID == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *RoleDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.Role{}).Error
}

func (dao *RoleDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.Role{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *RoleDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.Role{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *RoleDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&model.Role{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *RoleDao) Exsits(m *model.Role) bool {
	result := dao.db.First(m)
	return result.RowsAffected > 0
}
