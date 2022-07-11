package router

import (
	"github.com/labstack/echo/v4"

	"golang_blog/models"
	"golang_blog/mysql"
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
// @Router /article/{category_name} [post]
func createArticle(c echo.Context) error {
	contentName := c.Param("content_name")
	c.Set("content_name", contentName)
	var article models.Article
	if err := c.Bind(&article); err != nil {
		Render(c, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).CreateArticle(&article); err != nil {
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
// @Router /article/{content_name}/{article_name} [delete]
func deleteArticle(c echo.Context) error {
	contentName := c.Param("content_name")
	articleName := c.Param("article_name")
	c.Set("content_name", contentName)
	c.Set("article_name", articleName)
	var article models.Article
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).DeleteArticle(&article); err != nil {
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
// @Router /article/{content_name}/{article_name} [put]
func updateArticle(c echo.Context) error {
	contentName := c.Param("content_name")
	articleName := c.Param("article_name")
	c.Set("content_name", contentName)
	c.Set("article_name", articleName)
	var article models.Article
	if err := c.Bind(&article); err != nil {
		Render(c, err)
		return nil
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).UpdateArticle(&article); err != nil {
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
	var article models.Article
	var articleList []*models.Article
	var DB = mysql.DB
	articleList, err := workList.NewWorkList(c, DB).GetAllArticle(&article)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, articleList)
	return nil
}

// @Summary get article
// @Description get article
// @Tags Article
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /article/{user_id}/{category_id}/{article_id} [get]
func getArticle(c echo.Context) error {
	var articleParam = new(models.Article)
	if err := c.Bind(&articleParam); err != nil {
		Render(c, err)
		return err
	}
	var DB = mysql.DB
	article, err := workList.NewWorkList(c, DB).GetArticle(articleParam)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, article)
	return nil
}
