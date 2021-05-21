package chat

import (
	"blog/database"
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
var _ = chat.GET("/rooms", GetChatRoom)
var _ = chat.POST("/exit", ExitChatRoom)

func JoinChatRoom(c *gin.Context) {

	var reqData struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	if err := c.ShouldBind(&reqData); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "your post parameters are invalid",
		})
		return
	}

	cardId, err := c.Get("cardId")
	if !err {
		c.JSON(500, router.Response{
			State:   1,
			Message: "get cardId failed",
		})
		return
	}

	if err := joinRoom(reqData.Name, cardId.(string)); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "failed to into chat room: " + reqData.Name,
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

func GetChatRoom(c *gin.Context) {

	rooms, err := database.GetAllChatRoom()
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
		Data:    rooms,
	})
	return

}

func ExitChatRoom(c *gin.Context) {
	var reqData struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	if err := c.ShouldBind(&reqData); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "your post parameters are invalid",
		})
		return
	}

	cardId, err := c.Get("cardId")
	if !err {
		c.JSON(500, router.Response{
			State:   1,
			Message: "get cardId failed",
		})
		return
	}

	if err := exitRoom(reqData.Name, cardId.(string)); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "failed to quit chat room: " + reqData.Name,
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
	//TODO 首先需要检查他有没有加入该房间
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
		RoomName string `json:"room_name" binding:"required"`
		CardId   string `json:"cardId" binding:"required"`
	}

	if err := c.ShouldBind(&reqData); err != nil {
		log.Logger.Error(err.Error())
		c.JSON(200, router.Response{
			State:   1,
			Message: "your post parameters are invalid",
		})
		return
	}

	event, err := Consume(reqData.RoomName, reqData.CardId)
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
