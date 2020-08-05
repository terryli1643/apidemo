package dao

import (
	"errors"

	"github.com/terryli1643/apidemo/domain/model"
	"gorm.io/gorm"
)

type AdminDao struct {
	db *gorm.DB
}

func NewAdminDao(db *gorm.DB) *AdminDao {
	return &AdminDao{
		db: db,
	}
}

func (dao *AdminDao) Create(m *model.Admin) error {
	return dao.db.Create(m).Error
}

func (dao *AdminDao) Find(m *model.Admin) (result []model.Admin, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *AdminDao) FindOne(m *model.Admin) error {
	return dao.db.First(m, m).Error
}

func (dao *AdminDao) FindLast(m *model.Admin) error {
	return dao.db.Last(m, m).Error
}

func (dao *AdminDao) FindPage(m *model.Admin, rowbound model.RowBound, desc bool) (result []model.Admin, count int64, err error) {
	db := dao.db
	if desc {
		db = db.Order("id desc")
	}
	err = db.Model(m).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *AdminDao) Get(m *model.Admin) error {
	if m.ID == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *AdminDao) BatchGet(idbatch []int64) (result []model.Admin, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&model.Admin{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *AdminDao) GetForUpdate(m *model.Admin) error {
	if m.ID == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *AdminDao) Save(m *model.Admin) error {
	return dao.db.Save(m).Error
}

func (dao *AdminDao) Delete(m *model.Admin) error {
	if m.ID == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}

func (dao *AdminDao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.Admin{}).Error
}

func (dao *AdminDao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.Admin{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *AdminDao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.Admin{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *AdminDao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&model.Admin{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *AdminDao) Exsits(m *model.Admin) bool {
	result := dao.db.First(m)
	return result.RowsAffected > 0
}
