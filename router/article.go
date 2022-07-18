package router

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"golang_blog/models"
	"golang_blog/workList"
)

// @Summary create_article
// @Description create article
// @Tags Article
// @Security ApiKeyAuth
// @Param data body models.Article true "文章"
// @Accept json
// @Produce json
// @Success 200
// @Router /article [post]
func createArticle(c echo.Context) error {
	var article = new(models.Article)
	if err := c.Bind(&article); err != nil {
		Render(c, err)
		return err
	}
	err := workList.NewWorkList().CreateArticle(c, article)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary delete_article
// @Description delete article
// @Tags Article
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /article/{article_id} [delete]
func deleteArticle(c echo.Context) error {
	articleIdStr := c.Param("article_id")
	articleId, err := strconv.Atoi(articleIdStr)
	if err != nil {
		Render(c, err)
		return err
	}
	err = workList.NewWorkList().DeleteArticle(c, int32(articleId))
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary update_article
// @Description update article
// @Tags Article
// @Security ApiKeyAuth
// @Param data body models.Article true "文章"
// @Accept json
// @Produce json
// @Success 200
// @Router /article/{article_id} [put]
func updateArticle(c echo.Context) error {
	articleIdStr := c.Param("article_id")
	articleId, err := strconv.Atoi(articleIdStr)
	if err != nil {
		Render(c, err)
		return err
	}
	var article = new(models.Article)
	if err = c.Bind(&article); err != nil {
		Render(c, err)
		return err
	}
	err = workList.NewWorkList().UpdateArticle(c, int32(articleId), article)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary get_all_article
// @Description get all article
// @Tags Article
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /article [get]
func getAllArticle(c echo.Context) error {
	articles, err := workList.NewWorkList().GetAllArticle(c)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, articles)
	return nil
}

// @Summary get article
// @Description get article
// @Tags Article
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /article/{article_id} [get]
func getArticle(c echo.Context) error {
	articleIdStr := c.Param("article_id")
	articleId, err := strconv.Atoi(articleIdStr)
	if err != nil {
		Render(c, err)
		return err
	}
	articles, err := workList.NewWorkList().GetArticle(c, int32(articleId))
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, articles)
	return nil
}
