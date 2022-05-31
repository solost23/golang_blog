package router

import (
	"github.com/labstack/echo/v4"

	"golang_blog/model"
	"golang_blog/mysql"
	"golang_blog/workList"
)

// @Summary get_all_comment
// @Description get all comment
// @Tags Comment
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/{article_id} [get]
func getAllComment(c echo.Context) error {
	articleID := c.Param("article_id")
	c.Set("article_id", articleID)
	var comment model.Comment
	var commentList []*model.Comment
	var DB = mysql.DB
	commentList, err := workList.NewWorkList(c, DB).GetAllComment(&comment)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, commentList)
	return nil
}

// @Summary create_comment
// @Description create comment
// @Tags Comment
// @Security ApiKeyAuth
// @Param data body model.Comment true "评论"
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/{user_name}/{article_id}/{parent_id} [post]
func createComment(c echo.Context) error {
	userName := c.Param("user_name")
	articleID := c.Param("article_id")
	parentID := c.Param("parent_id")
	c.Set("user_name", userName)
	c.Set("article_id", articleID)
	c.Set("parent_id", parentID)
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		Render(c, err)
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).CreateComment(&comment); err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary delete comment
// @Description delete comment
// @Tags Comment
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /comment/{comment_id} [delete]
func deleteComment(c echo.Context) error {
	userName := c.Param("user_name")
	commentID := c.Param("comment_id")
	c.Set("user_name", userName)
	c.Set("comment_id", commentID)
	var comment = new(model.Comment)
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).DeleteComment(comment); err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}
