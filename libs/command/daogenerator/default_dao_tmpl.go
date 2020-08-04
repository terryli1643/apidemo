package main

var defaultDaoTmpl = `package dao
import (
	"errors"
	"time"

	"{{.ModelPkg}}"
	"github.com/shopspring/decimal"
	"gitlab.99safe.org/Shadow/shadow-framework/orm/jinzhu/gorm"
)

type {{.ModelName}}Dao struct {
	db *gorm.DB
}

func New{{.ModelName}}Dao(db *gorm.DB) *{{.ModelName}}Dao {
	return &{{.ModelName}}Dao{
		db: db,
	}
}

func (dao *{{.ModelName}}Dao) Create(m *model.{{.ModelName}}) error {
	return dao.db.Create(m).Error
}

func (dao *{{.ModelName}}Dao) Find(m *model.{{.ModelName}}) (result []model.{{.ModelName}}, err error) {
	err = dao.db.Find(&result, m).Error
	return
}

func (dao *{{.ModelName}}Dao) FindOne(m *model.{{.ModelName}}) error {
	return dao.db.First(m, m).Error
}

func (dao *{{.ModelName}}Dao) FindLast(m *model.{{.ModelName}}) error {
	return dao.db.Last(m, m).Error
}

func (dao *{{.ModelName}}Dao) FindPage(m *model.{{.ModelName}}, rowbound model.RowBound, desc bool)  (result []model.{{.ModelName}}, count int, err error)  {
	db := dao.db
	if desc {
		db = db.Order("id desc")
	}
	err = db.Model(m).Count(&count).Limit(rowbound.Limit).Offset(rowbound.Offset).Find(&result, m).Error
	return
}

func (dao *{{.ModelName}}Dao) Get(m *model.{{.ModelName}}) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Find(m).Error
}

func (dao *{{.ModelName}}Dao) BatchGet(idbatch []int64) (result []model.{{.ModelName}}, err error) {
	if len(idbatch) == 0 {
		return nil, errors.New("id is nil")
	}
	err = dao.db.Model(&model.{{.ModelName}}{}).Where("ID in (?)", idbatch).Find(&result).Error
	return
}

func (dao *{{.ModelName}}Dao) GetForUpdate(m *model.{{.ModelName}}) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Set("gorm:query_option", "FOR UPDATE").Find(m).Error
}

func (dao *{{.ModelName}}Dao) Save(m *model.{{.ModelName}}) error {
	return dao.db.Save(m).Error
}

func (dao *{{.ModelName}}Dao) Delete(m *model.{{.ModelName}}) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	{{if CheckField .Columns "deleted_at"}}return dao.db.Unscoped().Delete(m).Error
	{{else}}return dao.db.Delete(m).Error{{end}}
}

{{if CheckField .Columns "deleted_at"}}func (dao *{{.ModelName}}Dao) SoftDelete(m *model.{{.ModelName}}) error {
	if dao.db.NewRecord(m) {
		return errors.New("id is nil")
	}
	return dao.db.Delete(m).Error
}{{end}}

func (dao *{{.ModelName}}Dao) BatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	{{if CheckField .Columns "deleted_at"}}return dao.db.Unscoped().Where("ID in (?)", idbatch).Delete(&model.{{.ModelName}}{}).Error
	{{else}}return dao.db.Where("ID in (?)", idbatch).Delete(&model.{{.ModelName}}{}).Error{{end}}
}

{{if CheckField .Columns "deleted_at"}}func (dao *{{.ModelName}}Dao) SoftBatchDelete(idbatch []int64) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Where("ID in (?)", idbatch).Delete(&model.{{.ModelName}}{}).Error
}{{end}}

func (dao *{{.ModelName}}Dao) Updates(id int64, attrs map[string]interface{}) error {
	return dao.db.Model(&model.{{.ModelName}}{}).Where("ID = ?", id).Updates(attrs).Error
}

func (dao *{{.ModelName}}Dao) Update(id int64, attr string, value interface{}) error {
	return dao.db.Model(&model.{{.ModelName}}{}).Where("ID = ?", id).Update(attr, value).Error
}

func (dao *{{.ModelName}}Dao) BatchUpdaterAttrs(idbatch []int64, attrs map[string]interface{}) error {
	if len(idbatch) == 0 {
		return errors.New("id is nil")
	}
	return dao.db.Model(&model.{{.ModelName}}{}).Where("ID in (?)", idbatch).Updates(attrs).Error
}

func (dao *{{.ModelName}}Dao) Found(m *model.{{.ModelName}}) bool {
	find := dao.db.First(m, m).RecordNotFound()
	return !find
}
`
