package router

import (
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
func createCategory(c echo.Context) error {
	//fmt.Println(c.Get("token"))
	var category models.Category
	if err := c.Bind(&category); err != nil {
		Render(c, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).CreateCategory(&category); err != nil {
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
func deleteCategory(c echo.Context) error {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		Render(c, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).DeleteCategory(&category); err != nil {
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
func updateCategory(c echo.Context) error {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		Render(c, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).UpdateCategory(&category); err != nil {
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
func getAllCategory(c echo.Context) error {
	var content models.Category
	var contentList []*models.Category
	var DB = mysql.DB
	contentList, err := workList.NewWorkList(c, DB).GetAllCategory(&content)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, contentList)
	return nil
}
