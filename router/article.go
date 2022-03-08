package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"golang_blog/model"
	"golang_blog/mysql"
	"golang_blog/workList"
)

// PingExample godoc
// @Summary ping article
// @Schemes
// @Description create a article
// @Tags Article
// @Accept json
// @Produce json
// @Success 200
// @Router /article/{content_name} [get]
func createArticle(c echo.Context) error {
	contentName := c.Param("content_name")
	c.Set("content_name", contentName)
	var article model.Article
	if err := c.Bind(&article); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).CreateArticle(&article); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, "article create success")
	return nil
}

// PingExample godoc
// @Summary ping article
// @Schemes
// @Description delete a article
// @Tags Article
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
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, "article delete success")
	return nil
}

// PingExample godoc
// @Summary ping article
// @Schemes
// @Description create a article
// @Tags Article
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
		c.JSON(http.StatusBadRequest, err)
		return nil
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).UpdateArticle(&article); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, "article update success")
	return nil
}

// PingExample godoc
// @Summary ping article
// @Schemes
// @Description get all article
// @Tags Article
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
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, articleList)
	return nil
}

// PingExample godoc
// @Summary ping article
// @Schemes
// @Description get a article
// @Tags Article
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
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, article)
	return nil
}
