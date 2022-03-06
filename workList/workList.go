package workList

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type WorkList struct {
	conn *gorm.DB
	ctx  echo.Context
}

func NewWorkList(c echo.Context, conn *gorm.DB) *WorkList {
	return &WorkList{
		conn: conn,
		ctx:  c,
	}
}
