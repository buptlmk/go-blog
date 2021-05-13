package comment

import (
	"blog/database"
	"blog/log"
	"blog/middleware"
	"blog/router"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

var comment = router.App.Group("/comment")

var _ = comment.Use(middleware.Authorized)

var _ = comment.POST("/:id", AddComment)
var _ = comment.GET("/:id", GetComments)

func AddComment(c *gin.Context) {

	var comment = utils.Comment{}
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}

	err := database.AddComment(comment.UserId, comment.ArticleId, comment.Content)
	if err != nil {

		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, router.Response{
		State:   0,
		Message: "success",
	})

	return

}

func GetComments(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "article can't find",
		})
		return
	}

	comments, err := database.GetComments(int64(articleId))
	if err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}
	log.Logger.Info(comments)
	c.JSON(200, router.Response{
		State:   0,
		Message: "success",
		Data:    comments,
	})
}
