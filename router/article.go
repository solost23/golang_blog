package router

import (
	"github.com/labstack/echo/v4"

	"golang_blog/model"
	"golang_blog/mysql"
	"golang_blog/workList"
)

// @Summary create_article
// @Description create article
// @Tags Article
// @Security ApiKeyAuth
// @Param data body model.Article true "文章"
// @Accept json
// @Produce json
// @Success 200
// @Router /article/{content_name} [get]
func createArticle(c echo.Context) error {
	contentName := c.Param("content_name")
	c.Set("content_name", contentName)
	var article model.Article
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
	var article model.Article
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
// @Param data body model.Article true "文章"
// @Accept json
// @Produce json
// @Success 200
// @Router /article/{content_name}/{article_name} [put]
func updateArticle(c echo.Context) error {
	contentName := c.Param("content_name")
	articleName := c.Param("article_name")
	c.Set("content_name", contentName)
	c.Set("article_name", articleName)
	var article model.Article
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
	var article model.Article
	var articleList []*model.Article
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
// @Router /article/{user_name}/{content_name}/{article_name} [get]
func getArticle(c echo.Context) error {
	userName := c.Param("user_name")
	contentName := c.Param("content_name")
	articleName := c.Param("article_name")
	c.Set("user_name", userName)
	c.Set("content_name", contentName)
	c.Set("article_name", articleName)
	var article model.Article
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).GetArticle(&article); err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, article)
	return nil
}
