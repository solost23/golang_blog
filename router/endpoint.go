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
	group.POST("register", reg)
	group.POST("login", login)
}

func RegisterAuth(group *echo.Group) {
	group.Use(jwt.JWTAuth)
	RegisterUser(group)
	RegisterArticle(group)
	RegisterCategory(group)
	RegisterComment(group)

	group.Use(role.AuthCheckRole)
	RegisterRole(group)
}

func RegisterUser(group *echo.Group) {
	user := group.Group("user")
	{
		// 修改用户信息
		user.PUT(":user_id", updateUser)
		// 删除用户信息
		user.DELETE(":user_id", deleteUser)
	}
}

func RegisterArticle(group *echo.Group) {
	article := group.Group("article")
	{
		article.GET("", getAllArticle)
		article.GET(":article_id", getArticle)
		article.POST("", createArticle)
		article.DELETE(":article_id", deleteArticle)
		article.PUT(":article_id", updateArticle)
	}
}

func RegisterCategory(group *echo.Group) {
	content := group.Group("category")
	{
		content.GET("", getAllCategory)
		content.POST("", createCategory)
		content.DELETE(":category_id", deleteCategory)
		content.PUT(":category_id", updateCategory)
	}
}

func RegisterComment(group *echo.Group) {
	comment := group.Group("comment")
	{
		comment.GET(":article_id", getAllComment)
		comment.POST("", createComment)
		comment.DELETE(":comment_id", deleteComment)
	}
}

func RegisterRole(group *echo.Group) {
	role := group.Group("role")
	{
		role.POST("", addRoleAuth)
		role.DELETE(":role_id", deleteRoleAuth)
		role.GET("", getAllRoleAuth)
		role.GET(":role_name", getRoleAuth)
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
