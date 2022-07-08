package models

import (
	"errors"

	"gorm.io/gorm"
)

type CasbinModel struct {
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

func (t *CasbinModel) Insert(data interface{}) (err error) {
	return DBCasbin.Table(t.TableName()).Create(&t).Error
}

func (t *CasbinModel) Delete(query interface{}, args ...interface{}) (err error) {
	return DBCasbin.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *CasbinModel) Save(data interface{}, query interface{}, args ...interface{}) (err error) {
	return nil
}

func (t *CasbinModel) WhereOne(query interface{}, args ...interface{}) (casbinModel interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).First(&casbinModel).Error
	if err != nil {
		return nil, err
	}
	return casbinModel, nil
}

func (t *CasbinModel) WhereAll(query interface{}, args ...interface{}) (casbinModels interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).Find(&casbinModels).Error
	if err != nil {
		return nil, err
	}
	return casbinModels, nil
}

func (t *CasbinModel) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (casbinModels interface{}, count int64, err error) {
	offset := (params.Page - 1) * params.PageSize
	query := DB.Table(t.TableName()).Where(conditions, args...)

	err = query.Offset(offset).Limit(params.PageSize).Find(&casbinModels).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	err = DB.Table(t.TableName()).Where(conditions, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return casbinModels, count, nil
}

//func (c *CasbinModel) FindByRoleName(roleName string) ([]*CasbinModel, error) {
//	var res []*CasbinModel
//	if err := DBCasbin.Table(c.TableName()).Where("v0=?", roleName).Find(&res).Error; err != nil {
//		return res, err
//	}
//	return res, nil
//}
//
//func (c *CasbinModel) Find() ([]*CasbinModel, error) {
//	var res []*CasbinModel
//	if err := DBCasbin.Table(c.TableName()).Find(&res).Error; err != nil {
//		return res, err
//	}
//	return res, nil
//}
//
//func (c *CasbinModel) FindByRolePathMethod(roleName, path, method string) error {
//	if err := DBCasbin.Table(c.TableName()).Where("v0=? AND v1=? AND v2=?", roleName, path, method).First(c).Error; err != nil {
//		return err
//	}
//	return nil
//}
