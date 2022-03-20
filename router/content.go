package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang_blog/model"
	"golang_blog/mysql"
	"golang_blog/workList"
	"net/http"
)

// PingExample godoc
// @Summary ping content
// @Schemes
// @Description create a content
// @Tags Content
// @Accept json
// @Produce json
// @Success 200
// @Router /content [post]
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
	c.JSON(http.StatusOK, content)
	return nil
}

// PingExample godoc
// @Summary ping content
// @Schemes
// @Description delete a content
// @Tags Content
// @Accept json
// @Produce json
// @Success 200
// @Router /content [delete]
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

// PingExample godoc
// @Summary ping content
// @Schemes
// @Description update a content
// @Tags Content
// @Accept json
// @Produce json
// @Success 200
// @Router /content [put]
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
	c.JSON(http.StatusOK, content)
	return nil
}

// PingExample godoc
// @Summary ping content
// @Schemes
// @Description get all content
// @Tags Content
// @Accept json
// @Produce json
// @Success 200
// @Router /content [get]
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
