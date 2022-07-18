package router

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"golang_blog/models"
	"golang_blog/workList"
)

// @Summary create roleAuth
// @Description create roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Param data body models.CasbinModel true "角色权限"
// @Accept json
// @Produce json
// @Success 200
// @Router /role [post]
func addRoleAuth(c echo.Context) error {
	var casbinModel = new(models.CasbinModel)
	if err := c.Bind(casbinModel); err != nil {
		Render(c, err)
		return err
	}
	err := workList.NewWorkList().AddRoleAuth(c, casbinModel)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary delete roleAuth
// @Description delete roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /role/{role_id} [delete]
func deleteRoleAuth(c echo.Context) error {
	roleIdStr := c.Param("role_id")
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		Render(c, err)
		return err
	}
	var casbinModel = new(models.CasbinModel)
	if err = c.Bind(casbinModel); err != nil {
		Render(c, err)
		return err
	}
	err = workList.NewWorkList().DeleteRoleAuth(c, int32(roleId))
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary get_all_roleAuth
// @Description get all roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /role [get]
func getAllRoleAuth(c echo.Context) error {
	casbinModelList, err := workList.NewWorkList().GetAllRoleAuth(c)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, casbinModelList)
	return nil
}

// @Summary get_roleAuth
// @Description get roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /role/{role_name} [get]
func getRoleAuth(c echo.Context) error {
	roleName := c.Param("role_name")
	casbinModelList, err := workList.NewWorkList().GetRoleAuth(c, roleName)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, casbinModelList)
	return nil
}
