package router

import (
	"github.com/labstack/echo/v4"

	"golang_blog/model"
	"golang_blog/mysql"
	"golang_blog/workList"
)

// @Summary create roleAuth
// @Description create roleAuth
// @Tags Role
// @Security ApiKeyAuth
// @Param data body model.CasbinModel true "角色权限"
// @Accept json
// @Produce json
// @Success 200
// @Router /role [post]
func addRoleAuth(c echo.Context) error {
	var casbinModel = new(model.CasbinModel)
	if err := c.Bind(casbinModel); err != nil {
		Render(c, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).AddRoleAuth(casbinModel); err != nil {
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
// @Router /role [delete]
func deleteRoleAuth(c echo.Context) error {
	var casbinModel = new(model.CasbinModel)
	if err := c.Bind(casbinModel); err != nil {
		Render(c, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).DeleteRoleAuth(casbinModel); err != nil {
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
	var casbinModel = new(model.CasbinModel)
	var DB = mysql.DB
	casbinModelList, err := workList.NewWorkList(c, DB).GetAllRoleAuth(casbinModel)
	if err != nil {
		Render(c, err)
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
	c.Set("role_name", roleName)
	var casbinModel = new(model.CasbinModel)
	var DB = mysql.DB
	casbinModelList, err := workList.NewWorkList(c, DB).GetRoleAuth(casbinModel)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, casbinModelList)
	return nil
}
