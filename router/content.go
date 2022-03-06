package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"jwt-go/model"
	"jwt-go/mysql"
	"jwt-go/workList"
	"net/http"
)

func createContent(c echo.Context) error {
	//fmt.Println(c.Get("token"))
	var content model.Content
	if err := c.Bind(&content); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).CreateContent(&content); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, "content create success")
	return nil
}

func deleteContent(c echo.Context) error {
	var content model.Content
	if err := c.Bind(&content); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).DeleteContent(&content); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, "content delete success")
	return nil
}

func updateContent(c echo.Context) error {
	var content model.Content
	if err := c.Bind(&content); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	fmt.Println(content)
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).UpdateContent(&content); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, "content update success")
	return nil
}

func getAllContent(c echo.Context) error {
	var content model.Content
	var contentList []*model.Content
	var DB = mysql.DB
	contentList, err := workList.NewWorkList(c, DB).GetAllContent(&content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, contentList)
		return err
	}
	c.JSON(http.StatusOK, contentList)
	return nil
}
