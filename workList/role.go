package workList

import (
	"errors"
	"gorm.io/gorm"

	"golang_blog/model"
)

func (w *WorkList) AddRoleAuth(casbinModel *model.CasbinModel) error {
	// 先查询本条数据是否存在
	// 若不存在，则插入
	if err := casbinModel.FindByRolePathMethod(casbinModel.RoleName, casbinModel.Path, casbinModel.Method); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if err := casbinModel.Create(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteRoleAuth(casbinModel *model.CasbinModel) error {
	// 先查询本条数据是否存在
	// 若存在，则删除
	if err := casbinModel.FindByRolePathMethod(casbinModel.RoleName, casbinModel.Path, casbinModel.Method); err != nil {
		return err
	}
	if err := casbinModel.Delete(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetAllRoleAuth(casbinModel *model.CasbinModel) ([]*model.CasbinModel, error) {
	// 直接获取所有
	casbinModelList, err := casbinModel.Find()
	if err != nil {
		return casbinModelList, err
	}
	return casbinModelList, nil
}

func (w *WorkList) GetRoleAuth(casbinModel *model.CasbinModel) ([]*model.CasbinModel, error) {
	// 直接查找，若找不到，返回错误
	roleName := w.ctx.Get("role_name").(string)
	res, err := casbinModel.FindByRoleName(roleName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, err
		}
		return res, err
	}
	return res, nil
}
