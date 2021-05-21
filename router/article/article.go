package article

import (
	"blog/database"
	"blog/log"
	"blog/router"
	"github.com/gin-gonic/gin"
	"strconv"
)

var _ = router.App.GET("/article/:id", GetArticle)
var _ = router.App.GET("/article/recent/:number", GetRecentArticle)

var article = router.App.Group("/article")
var _ = article.GET("/star/:id", StarArticle)

//var _ = router.App.GET("/article/star/:id", middleware.Authorized, StarArticle)
var _ = router.App.POST("/article/edit", AddArticle)

// 考虑写完文章，令牌过期，用户无法提交，所以提交这一步骤，不再走权限认证
func AddArticle(c *gin.Context) {

	var reqData struct {
		Id      int    `json:"id"`
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	err := c.ShouldBind(&reqData)
	if err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}
	log.Logger.Info(reqData)
	err = database.AddArticle(reqData.Id, reqData.Title, reqData.Content)
	if err != nil {
		log.Logger.Error(err.Error())
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

func StarArticle(c *gin.Context) {

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "article id is wrong",
		})
		return
	}
	err = database.StarArticle(int64(id))
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

func GetArticle(c *gin.Context) {

	id := c.Param("id")
	articleId, err := strconv.Atoi(id)
	if err != nil {
		log.Logger.Info(err.Error())
		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}
	article, err := database.GetArticle(int64(articleId))
	if err != nil {
		log.Logger.Info(err.Error())
		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}

	err = database.AddArticleVisitedNumber(int64(articleId), 1)
	if err != nil {
		log.Logger.Error(err.Error())
	}

	c.JSON(200, router.Response{
		State:   0,
		Data:    article,
		Message: "success",
	})

}

func GetRecentArticle(c *gin.Context) {

	log.Logger.Info(c.Request.URL.Path)
	param := c.Param("number")
	number, err := strconv.Atoi(param)
	if err != nil {
		log.Logger.Info(err.Error())
		c.JSON(500, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}
	articles, err := database.GetRecentArticle(number, 0)
	if err != nil {
		log.Logger.Info(err.Error())
		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, router.Response{
		State:   0,
		Data:    articles,
		Message: "success",
	})

}
