package workList

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"golang_blog/models"
)

func (w *WorkList) AddRoleAuth(casbinModelParam *models.CasbinModel) error {
	query := []string{"v0 = ?", "v1 = ?", "v2 = ?"}
	args := []interface{}{casbinModelParam.RoleName, casbinModelParam.Path, casbinModelParam.Method}
	_, err := models.NewCasbinModel().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err = models.NewCasbinModel().Insert(casbinModelParam); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteRoleAuth(casbinModelParam *models.CasbinModel) error {
	// 先查询本条数据是否存在
	// 若存在，则删除
	query := []string{"v0 = ?", "v1 = ?", "v2 = ?"}
	args := []interface{}{casbinModelParam.RoleName, casbinModelParam.Path, casbinModelParam.Method}
	_, err := models.NewCasbinModel().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err = models.NewCasbinModel().Delete(strings.Join(query, " AND "), args...); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetAllRoleAuth(casbinModelParam *models.CasbinModel) ([]*models.CasbinModel, error) {
	// 直接获取所有
	casbinModelList, err := models.NewCasbinModel().WhereAll(nil, nil)
	if err != nil {
		return nil, err
	}
	return casbinModelList.([]*models.CasbinModel), nil
}

func (w *WorkList) GetRoleAuth(casbinModelParam *models.CasbinModel) ([]*models.CasbinModel, error) {
	// 直接查找，若找不到，返回错误
	roleName := w.ctx.Get("role_name").(string)
	query := []string{"v0 = ?"}
	args := []interface{}{roleName}
	casbinModels, err := models.NewCasbinModel().WhereAll(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return casbinModels.([]*models.CasbinModel), nil
}
