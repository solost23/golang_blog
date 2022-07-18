package router

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"golang_blog/models"
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
	articleIdStr := c.Param("article_id")
	articleId, err := strconv.Atoi(articleIdStr)
	if err != nil {
		Render(c, err)
		return err
	}
	comments, err := workList.NewWorkList().GetAllComment(c, int32(articleId))
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil, comments)
	return nil
}

// @Summary create_comment
// @Description create comment
// @Tags Comment
// @Security ApiKeyAuth
// @Param data body models.Comment true "评论"
// @Accept json
// @Produce json
// @Success 200
// @Router /comment [post]
func createComment(c echo.Context) error {
	var comment = new(models.Comment)
	if err := c.Bind(&comment); err != nil {
		Render(c, err)
		return err
	}
	err := workList.NewWorkList().CreateComment(c, comment)
	if err != nil {
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
	commentIdStr := c.Param("comment_id")
	commentId, err := strconv.Atoi(commentIdStr)
	if err != nil {
		Render(c, err)
		return err
	}
	err = workList.NewWorkList().DeleteComment(c, int32(commentId))
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}
