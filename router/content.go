package router

import (
	"fmt"
	"golang_blog/models"
	"golang_blog/mysql"
	"golang_blog/workList"

	"github.com/labstack/echo/v4"
)

// @Summary create content
// @Description create a content
// @Tags Content
// @Security ApiKeyAuth
// @Param data body models.Content true "分类"
// @Accept json
// @Produce json
// @Success 200
// @Router /content [post]
func createContent(c echo.Context) error {
	//fmt.Println(c.Get("token"))
	var content models.Content
	if err := c.Bind(&content); err != nil {
		Render(c, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).CreateContent(&content); err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary delete content
// @Description delete content
// @Tags Content
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /content [delete]
func deleteContent(c echo.Context) error {
	var content models.Content
	if err := c.Bind(&content); err != nil {
		Render(c, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).DeleteContent(&content); err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary update content
// @Description update content
// @Tags Content
// @Security ApiKeyAuth
// @Param data body models.Content true "分类"
// @Accept json
// @Produce json
// @Success 200
// @Router /content [put]
func updateContent(c echo.Context) error {
	var content models.Content
	if err := c.Bind(&content); err != nil {
		Render(c, err)
		return err
	}
	fmt.Println(content)
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).UpdateContent(&content); err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary get_all_content
// @Description get all content
// @Tags Content
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /content [get]
func getAllContent(c echo.Context) error {
	var content models.Content
	var contentList []*models.Content
	var DB = mysql.DB
	contentList, err := workList.NewWorkList(c, DB).GetAllContent(&content)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, contentList)
	return nil
}
