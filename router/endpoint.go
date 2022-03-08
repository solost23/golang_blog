package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	_ "golang_blog/docs" // 一定要导入docs，否则会报内部错误
	"golang_blog/middleware/jwt"
)

func RegisterNoAuth(group *echo.Group) {
	user := group.Group("/user")
	{
		user.POST("/register", reg)
		user.POST("/login", login)
	}
}

func RegisterAuth(group *echo.Group) {
	group.Use(jwt.JWTAuth)

	RegisterArticle(group)
	RegisterContent(group)
}

func RegisterArticle(group *echo.Group) {
	article := group.Group("/article")
	{
		article.GET("", getAllArticle)
		// 看单篇文章可能要看别人的，所以要通过路径参数传用户名
		article.GET("/:user_name/:content_name/:article_name", getArticle)
		article.POST("/:content_name", createArticle)
		article.DELETE("/:content_name/:article_name", deleteArticle)
		article.PUT("/:content_name/:article_name", updateArticle)
	}
}

func RegisterContent(group *echo.Group) {
	content := group.Group("/content")
	{
		content.GET("", getAllContent)
		content.POST("", createContent)
		content.DELETE("", deleteContent)
		content.PUT("", updateContent)
	}
}

func Register() *echo.Echo {
	router := echo.New()
	// , middleware.Recover()
	router.Use(middleware.Logger())

	group := router.Group("")
	RegisterNoAuth(group)
	RegisterAuth(group)

	// swagger
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	return router
}
