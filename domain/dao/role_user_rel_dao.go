package dao

import (
	"errors"

	"github.com/terryli1643/apidemo/domain/model"
	"gorm.io/gorm"
)

type RoleUserRelDao struct {
	db *gorm.DB
}

func NewRoleUserRelDao(db *gorm.DB) *RoleUserRelDao {
	return &RoleUserRelDao{
		db: db,
	}
}

func (dao *RoleUserRelDao) Create(m *model.RoleUserRel) error {
	return dao.db.Create(m).Error
}

func (dao *RoleUserRelDao) Find(m *model.RoleUserRel) (result []model.RoleUserRel, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *RoleUserRelDao) FindOne(m *model.RoleUserRel) error {
	return dao.db.First(m, m).Error
}

func (dao *RoleUserRelDao) FindLast(m *model.RoleUserRel) error {
	return dao.db.Last(m, m).Error
}

func (dao *RoleUserRelDao) FindPage(m *model.RoleUserRel, rowbound model.RowBound, desc bool) (result []model.RoleUserRel, count int64, err error) {
	db := dao.db
	if desc {
		db = db.Order("id desc")
	}
	err = db.Model(m).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *RoleUserRelDao) Get(m *model.RoleUserRel) error {
	if m.ID == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *RoleUserRelDao) BatchGet(idbatch []int64) (result []model.RoleUserRel, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&model.RoleUserRel{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *RoleUserRelDao) GetForUpdate(m *model.RoleUserRel) error {
	if m.ID == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *RoleUserRelDao) Save(m *model.RoleUserRel) error {
	return dao.db.Save(m).Error
}

func (dao *RoleUserRelDao) Delete(m *model.RoleUserRel) error {
	if m.ID == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *RoleUserRelDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.RoleUserRel{}).Error
}

func (dao *RoleUserRelDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.RoleUserRel{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *RoleUserRelDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.RoleUserRel{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *RoleUserRelDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&model.RoleUserRel{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *RoleUserRelDao) Exsits(m *model.RoleUserRel) bool {
	result := dao.db.First(m)
	return result.RowsAffected > 0
}
