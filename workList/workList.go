package workList

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	NONE    = "NONE"
	INSERT  = "INSERT"
	DELETE  = "DELETE"
	UPDATE  = "UPDATE"
	SUCCESS = "SUCCESS"
	FAILED  = "FAILED"
	CONTENT = "CONTENT"
	USER    = "USER"
	ARTICLE = "ARTICLE"
	COMMENT = "COMMENT"
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
