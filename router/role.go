package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"golang_blog/model"
	"golang_blog/mysql"
	"golang_blog/workList"
)

// PingExample godoc
// @Summary ping role
// @Schemes
// @Description add roleAuth
// @Tags Role
// @Accept json
// @Produce json
// @Success 200
// @Router /role [post]
func addRoleAuth(c echo.Context) error {
	var casbinModel = new(model.CasbinModel)
	if err := c.Bind(casbinModel); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).AddRoleAuth(casbinModel); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, casbinModel)
	return nil
}

// PingExample godoc
// @Summary ping role
// @Schemes
// @Description delete roleAuth
// @Tags Role
// @Accept json
// @Produce json
// @Success 200
// @Router /role [delete]
func deleteRoleAuth(c echo.Context) error {
	var casbinModel = new(model.CasbinModel)
	if err := c.Bind(casbinModel); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).DeleteRoleAuth(casbinModel); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, "delete roleAuth success")
	return nil
}

// PingExample godoc
// @Summary ping role
// @Schemes
// @Description get all roleAuth
// @Tags Role
// @Accept json
// @Produce json
// @Success 200
// @Router /role [get]
func getAllRoleAuth(c echo.Context) error {
	var casbinModel = new(model.CasbinModel)
	var DB = mysql.DB
	casbinModelList, err := workList.NewWorkList(c, DB).GetAllRoleAuth(casbinModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, casbinModelList)
	return nil
}

// PingExample godoc
// @Summary ping role
// @Schemes
// @Description get roleAuth
// @Tags Role
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
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, casbinModelList)
	return nil
}
