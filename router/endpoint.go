package router

import (
	"golang_blog/middleware/role"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"

	_ "golang_blog/docs" // 一定要导入docs，否则会报内部错误
	"golang_blog/middleware/jwt"
	"golang_blog/middleware/logger"
)

type ErrCode int

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

// 封装返回
type ApiResponse struct {
	// in: body
	Code    ErrCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 封装返回
func RenderJSON(ctx echo.Context, code ErrCode, message string, data ...interface{}) {
	// 如果切片中无数据，那么不封装data，否则封装data
	var res ApiResponse
	if len(data) >= 1 {
		res = ApiResponse{
			Code:    code,
			Message: message,
			Data:    data[0],
		}
	} else {
		res = ApiResponse{
			Code:    code,
			Message: message,
		}
	}
	ctx.JSON(http.StatusOK, res)
}

func Render(ctx echo.Context, err error, data ...interface{}) {
	if err != nil {
		RenderJSON(ctx, http.StatusOK, err.Error(), data...)
	} else {
		RenderJSON(ctx, http.StatusOK, "success", data...)
	}
}
