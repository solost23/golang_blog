package role

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/labstack/echo/v4"

	"golang_blog/config"
	"golang_blog/middleware/jwt"
)

func AuthCheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("claims").(*jwt.Claims)
		role := claims.Role
		mysqlConfig := config.GetMysqlConfig()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Ip)
		a := xormadapter.NewAdapter("mysql", dsn)
		e := casbin.NewEnforcer("config/rbac_model.conf", a)
		if err := e.LoadPolicy(); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return err
		}
		ok := e.Enforce(role, c.Request().URL.Path, c.Request().Method)
		if !ok {
			c.JSON(http.StatusInternalServerError, "AuthCheckRole Failed")
			return nil
		}
		next(c)
		return nil
	}
}
