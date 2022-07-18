package router

import (
	"golang_blog/models"
	"golang_blog/workList"
	"strconv"

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
// @Router /category [post]
func createCategory(c echo.Context) error {
	var category = new(models.Category)
	if err := c.Bind(&category); err != nil {
		Render(c, err)
		return err
	}
	err := workList.NewWorkList().CreateCategory(c, category)
	if err != nil {
		Render(c, err)
		return err
	}
	return nil
}

// @Summary delete content
// @Description delete content
// @Tags Content
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /category/{category_id} [delete]
func deleteCategory(c echo.Context) error {
	categoryIdStr := c.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		Render(c, err)
		return err
	}
	var category = new(models.Category)
	if err = c.Bind(&category); err != nil {
		Render(c, err)
		return err
	}
	err = workList.NewWorkList().DeleteCategory(c, int32(categoryId))
	if err != nil {
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
// @Router /category/{category_id} [put]
func updateCategory(c echo.Context) error {
	categoryIdStr := c.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		Render(c, err)
		return err
	}
	var category = new(models.Category)
	if err = c.Bind(&category); err != nil {
		Render(c, err)
		return err
	}
	err = workList.NewWorkList().UpdateCategory(c, int32(categoryId), category)
	if err != nil {
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
// @Router /category [get]
func getAllCategory(c echo.Context) error {
	categories, err := workList.NewWorkList().GetAllCategory(c)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, categories)
	return nil
}
