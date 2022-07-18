package models

import (
	"errors"

	"gorm.io/gorm"
)

type CasbinModel struct {
	ID       int32  `gorm:"primary_key"`
	Ptype    string `gorm:"column:p_type;default:p"`
	RoleName string `gorm:"column:v0" json:"role_name"`
	Path     string `gorm:"column:v1" json:"path"`
	Method   string `gorm:"column:v2" json:"method"`
}

func NewCasbinModel() Moder {
	return &CasbinModel{}
}

func (t *CasbinModel) TableName() string {
	return "casbin_rule"
}

func (t *CasbinModel) Insert() (err error) {
	return DBCasbin.Table(t.TableName()).Create(&t).Error
}

func (t *CasbinModel) Delete(query interface{}, args ...interface{}) (err error) {
	return DBCasbin.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *CasbinModel) Save(query interface{}, args ...interface{}) (err error) {
	return nil
}

func (t *CasbinModel) WhereOne(query interface{}, args ...interface{}) (interface{}, error) {
	var casbinModel = new(CasbinModel)
	err := DBCasbin.Table(t.TableName()).Where(query, args...).First(casbinModel).Error
	if err != nil {
		return casbinModel, err
	}
	return casbinModel, nil
}

func (t *CasbinModel) WhereAll(query interface{}, args ...interface{}) (interface{}, error) {
	var casbinModels []*CasbinModel
	err := DBCasbin.Table(t.TableName()).Where(query, args...).Find(&casbinModels).Error
	if err != nil {
		return casbinModels, err
	}
	return casbinModels, nil
}

func (t *CasbinModel) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (interface{}, int64, error) {
	var casbinModels = new(CasbinModel)
	var count int64
	var err error
	offset := (params.Page - 1) * params.PageSize
	query := DBCasbin.Table(t.TableName()).Where(conditions, args...)

	err = query.Offset(offset).Limit(params.PageSize).Find(casbinModels).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return casbinModels, 0, err
	}

	err = DBCasbin.Table(t.TableName()).Where(conditions, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return casbinModels, 0, err
	}
	return casbinModels, count, nil
}
