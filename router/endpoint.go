package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"golang_blog/middleware/role"

	_ "golang_blog/docs" // 一定要导入docs，否则会报内部错误
	"golang_blog/middleware/jwt"
	"golang_blog/middleware/logger"
)

func RegisterNoAuth(group *echo.Group) {
	group.POST("/register", reg)
	group.POST("/login", login)
}

func RegisterAuth(group *echo.Group) {
	group.Use(jwt.JWTAuth)
	RegisterUser(group)
	RegisterArticle(group)
	RegisterContent(group)
	RegisterComment(group)

	group.Use(role.AuthCheckRole)
	RegisterRole(group)
}

func RegisterUser(group *echo.Group) {
	user := group.Group("/user")
	{
		// 修改用户信息
		user.PUT("/:user_name", updateUser)
		// 删除用户信息
		user.DELETE("/:user_name", deleteUser)
	}
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

func RegisterComment(group *echo.Group) {
	comment := group.Group("/comment")
	{
		comment.GET("/:article_id", getAllComment)
		comment.POST("/:user_name/:article_id/:parent_id", createComment)
		comment.DELETE("/:user_name/:comment_id", deleteComment)
	}
}

func RegisterRole(group *echo.Group) {
	role := group.Group("/role")
	{
		role.POST("", addRoleAuth)
		role.DELETE("", deleteRoleAuth)
		role.GET("", getAllRoleAuth)
		role.GET("/:role_name", getRoleAuth)
	}
}

func Register() *echo.Echo {
	router := echo.New()
	router.Use(logger.Logger, middleware.Recover())
	group := router.Group("")
	RegisterNoAuth(group)
	RegisterAuth(group)

	// swagger
	router.GET("/swagger/*", echoSwagger.WrapHandler)
	return router
}
