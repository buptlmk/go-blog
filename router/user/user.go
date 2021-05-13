package user

import (
	"blog/database"
	"blog/log"
	"blog/middleware"
	"blog/router"
	"blog/utils"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

var _ = router.App.POST("/login", loginUser)
var _ = router.App.POST("/register", registerUser)

var user = router.App.Group("user")
var _ = user.Use(middleware.Authorized)
var _ = user.GET("/detail", getUserInfo)

func getUserInfo(c *gin.Context) {

	v, exist := c.Get("cardId")
	if !exist {
		c.JSON(200, router.Response{
			State:   1,
			Message: "can't get cardId",
		})
		return
	}
	user, err := database.GetUserInfo(v.(string))
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
		Data:    user,
	})
	return

}

func loginUser(c *gin.Context) {

	var reqData struct {
		CardId   string `json:"card_id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBind(&reqData); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}

	id, passwordHash, saltHash, name, err := database.GetUserPassword(reqData.CardId)
	if err != nil {
		c.JSON(500, router.Response{
			State:   1,
			Message: err.Error(),
		})
		return
	}

	hash256 := sha256.New()
	hash256.Write([]byte(reqData.Password + saltHash))
	t := hex.EncodeToString(hash256.Sum(nil))
	if t != passwordHash {
		c.JSON(200, router.Response{
			State:   1,
			Message: "your id or password is incorrect",
		})
		return
	}

	// 生成token

	token, err := utils.GenerateToken(id, reqData.CardId)
	if err != nil {
		c.JSON(200, router.Response{
			State:   0,
			Message: err.Error(),
		})
		return
	}
	log.Logger.Info(token, name)
	http.SetCookie(c.Writer, &http.Cookie{Name: "token", Value: token, MaxAge: 3 * 86400, Path: "/"})
	http.SetCookie(c.Writer, &http.Cookie{Name: "id", Value: strconv.FormatInt(id, 10), MaxAge: 3 * 86400, Path: "/"})
	http.SetCookie(c.Writer, &http.Cookie{Name: "name", Value: name, MaxAge: 3 * 86400, Path: "/"})
	c.JSON(200, router.Response{
		State:   0,
		Message: "success",
	})
	return
}

func registerUser(c *gin.Context) {

	var reqData struct {
		CardId   string `json:"card_id" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBind(&reqData); err != nil {
		c.JSON(200, router.Response{
			State:   1,
			Message: "your input is incorrect",
		})
		return
	}
	log.Logger.Info(reqData)

	// 随机
	rnd, _ := rand.Int(rand.Reader, big.NewInt(1<<63-1))
	hash256 := sha256.New()
	hash256.Write([]byte(rnd.String()))
	salt := hex.EncodeToString(hash256.Sum(nil))

	hash256.Reset()
	hash256.Write([]byte(reqData.Password + salt))
	passSaltHash := hex.EncodeToString(hash256.Sum(nil))

	if err := database.RegisterUser(reqData.CardId, reqData.Name, passSaltHash, salt, reqData.Phone, reqData.Email); err != nil {
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
