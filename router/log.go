package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"golang_blog/model"
	"golang_blog/mysql"
	"golang_blog/workList"
)

// PingExample godoc
// @Summary ping log
// @Schemes
// @Description get all log
// @Tags Log
// @Accept json
// @Produce json
// @Success 200
// @Router /log [get]
func getAllLog(c echo.Context) error {
	var DB = mysql.DB
	var log model.Log
	logList, err := workList.NewWorkList(c, DB).GetAllLog(&log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, logList)
	return nil
}

// PingExample godoc
// @Summary ping log
// @Schemes
// @Description delete a log
// @Tags Log
// @Accept json
// @Produce json
// @Success 200
// @Router /log/{log_id} [delete]
func deleteLog(c echo.Context) error {
	id := c.Param("log_id")
	c.Set("id", id)
	var DB = mysql.DB
	var log model.Log
	if err := workList.NewWorkList(c, DB).DeleteLog(&log); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, "delete log success")
	return nil
}
