package activity

import (
	"blog/database"
	"blog/database/redis"
	"blog/log"
	"blog/middleware"
	"blog/router"
	"github.com/gin-gonic/gin"
	"strconv"
)

var _ = router.App.GET("/activity", GetActivity)

var activity = router.App.Group("/activity")
var _ = activity.Use(middleware.Authorized)

var _ = activity.GET("/join/:id", JoinActivity)
var _ = activity.POST("/add", AddActivity)

func GetActivity(c *gin.Context) {

	activities, err := database.GetAllActivity()
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
		Data:    activities,
	})
	return
}

func JoinActivity(c *gin.Context) {

	param := c.Param("id")

	activityId, err := strconv.Atoi(param)

	if err != nil {
		c.JSON(500, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}
	log.Logger.Info(activityId)

	//err = database.JoinActivity(activityId)
	//if err!=nil{
	//	c.JSON(200,router.Response{
	//		State: 1,
	//		Message: err.Error(),
	//	})
	//	return
	//}

	// 订单信息的hash值先用用户id表示
	v, exist := c.Get("cardId")
	if !exist {
		c.JSON(500, router.Response{
			State:   1,
			Message: "not found cardId",
		})
		return
	}
	err = redis.PurchaseTicket(activityId, v.(string))
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

func AddActivity(c *gin.Context) {

	var reqData struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
		Time        string `json:"time" binding:"required"`
		Img         string `json:"img"`
		Total       int    `json:"total" binding:"required"`
		Res         int    `json:"res" binding:"required"`
	}

	if err := c.ShouldBind(&reqData); err != nil {
		log.Logger.Error(err.Error())
		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}
	id, err := database.AddActivity(reqData.Name, reqData.Description, reqData.Img, reqData.Time, reqData.Total, reqData.Res)
	if err != nil {
		log.Logger.Error(err.Error())
		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}
	err = redis.AddActivity(int(id), reqData.Name, reqData.Total, 0)
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
}
