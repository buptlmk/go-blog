package chat

import (
	"blog/log"
	"blog/middleware"
	"blog/router"
	"github.com/gin-gonic/gin"
	"time"
)

var chat = router.App.Group("/chat")

var _ = chat.Use(middleware.Authorized)

var _ = chat.POST("/join", JoinChatRoom)
var _ = chat.POST("/add/:id", AddChatRoom)
var _ = chat.POST("/push/:roomName", Publish)
var _ = chat.POST("/pull", GetMessage)

func JoinChatRoom(c *gin.Context) {

	var reqData struct {
		RoomName string `json:"room_name"`
		UserId   string `json:"user_id"`
	}

	if err := c.ShouldBind(&reqData); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "your post parameters are invalid",
		})
		return
	}

	if err := joinRoom(reqData.RoomName, reqData.UserId); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "failed to into chat room: " + reqData.RoomName,
		})
		return
	}
	c.JSON(200, router.Response{
		State:   0,
		Message: "success",
	})
	return

}

func AddChatRoom(c *gin.Context) {

	roomName := c.Param("id")
	err := addRoom(roomName)
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

func Publish(c *gin.Context) {

	roomName := c.Param("roomName")

	var event = Event{}
	if err := c.ShouldBind(&event); err != nil {
		log.Logger.Error(err.Error())
		c.JSON(200, router.Response{
			State:   1,
			Message: "your post parameters are invalid",
		})
		return
	}
	event.Timestamp = time.Now()
	log.Logger.Info(event)
	if err := publish(roomName, event); err != nil {
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

func GetMessage(c *gin.Context) {
	var reqData struct {
		RoomName string `json:"room_name"`
		UserId   string `json:"user_id"`
	}

	if err := c.ShouldBind(&reqData); err != nil {
		log.Logger.Error(err.Error())
		c.JSON(200, router.Response{
			State:   1,
			Message: "your post parameters are invalid",
		})
		return
	}

	event, err := Consume(reqData.RoomName, reqData.UserId)
	if err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}
	log.Logger.Info(event)
	c.JSON(200, router.Response{
		State:   0,
		Data:    event,
		Message: "success",
	})

	return

}
