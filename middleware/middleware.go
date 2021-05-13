package middleware

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func Authorized(c *gin.Context) {

	// 取token
	token := c.GetHeader("token")
	if token == "" {
		c.Abort()
		c.JSON(utils.StatusUnauthorized, gin.H{
			"state":   1,
			"message": "token is null",
			"data":    nil,
		})
		return
	}

	info, err := utils.ParseToken(token)
	if err != nil {
		c.Abort()
		c.JSON(utils.StatusUnauthorized, gin.H{
			"state":   1,
			"message": "failed to parse token",
			"data":    nil,
		})
		return
	}
	if info.Expire {
		c.Abort()
		c.JSON(utils.StatusUnauthorized, gin.H{
			"state":   1,
			"message": "expired",
			"data":    nil,
		})
		return
	}

	c.Set("id", info.ID)
	c.Set("cardId", info.CardID)
	return

}

func RateLimitBucket() func(c *gin.Context) {
	// 1s 1000个
	bucket := utils.NewBucket(1*time.Millisecond, 1000)

	return func(c *gin.Context) {

		if ok := bucket.WaitMaxDuration(1, 3*time.Second); !ok {
			c.Abort()
			c.JSON(utils.StatusServiceUnavailable, gin.H{
				"state":   1,
				"message": "current visit is too large",
				"data":    nil,
			})
			return
		}
		c.Next()
	}
}
