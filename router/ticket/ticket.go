package ticket

import (
	"blog/database"
	"blog/database/redis"
	"blog/middleware"
	"blog/router"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"math/big"
)

var ticket = router.App.Group("/ticket")
var _ = ticket.Use(middleware.Authorized)

var _ = ticket.POST("/new", NewTicket)

var _ = ticket.GET("/get", GetTicket)

func NewTicket(c *gin.Context) {

	var reqData struct {
		Name   string `json:"name"`
		Number int64  `json:"number"`
	}

	if err := c.ShouldBind(&reqData); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "信息错误",
		})
		return
	}

	redis.InitTicket(reqData.Number)
	err := database.SetTicket(reqData.Name, reqData.Number)
	if err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "信息错误",
		})
		return
	}

	c.JSON(200, router.Response{
		State:   0,
		Message: "success",
	})
	return
}

func GetTicket(c *gin.Context) {

	rnd, _ := rand.Int(rand.Reader, big.NewInt(1<<63-1))

	sha := sha256.New()
	sha.Write([]byte(rnd.String()))
	orderHash := hex.EncodeToString(sha.Sum(nil))

	if err := redis.Pur(orderHash); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "failed",
		})
		return
	}

	c.JSON(200, router.Response{
		State:   0,
		Message: "success",
	})

	return

}
