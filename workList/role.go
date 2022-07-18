package workList

import (
	"errors"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"

	"golang_blog/models"
)

func (w *WorkList) AddRoleAuth(_ echo.Context, casbinModelParam *models.CasbinModel) error {
	query := []string{"v0 = ?", "v1 = ?", "v2 = ?"}
	args := []interface{}{casbinModelParam.RoleName, casbinModelParam.Path, casbinModelParam.Method}
	_, err := models.NewCasbinModel().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err == nil {
		return errors.New("data is exist")
	}

	if err = casbinModelParam.Insert(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteRoleAuth(_ echo.Context, id int32) error {
	// 先查询本条数据是否存在
	// 若存在，则删除
	query := []string{"id = ?"}
	args := []interface{}{id}
	_, err := models.NewCasbinModel().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err = models.NewCasbinModel().Delete(strings.Join(query, " AND "), args...); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetAllRoleAuth(_ echo.Context) ([]*models.CasbinModel, error) {
	// 直接获取所有
	query := []string{"1 = ?"}
	args := []interface{}{1}
	casbinModelList, err := models.NewCasbinModel().WhereAll(strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	return casbinModelList.([]*models.CasbinModel), nil
}

func (w *WorkList) GetRoleAuth(_ echo.Context, roleName string) ([]*models.CasbinModel, error) {
	// 直接查找，若找不到，返回错误
	fmt.Println("roleName:", roleName)
	query := []string{"v0 = ?"}
	args := []interface{}{roleName}
	casbinModels, err := models.NewCasbinModel().WhereAll(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return casbinModels.([]*models.CasbinModel), nil
}
