package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"jwt-go/model"
	"jwt-go/mysql"
	"jwt-go/workList"
)

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
