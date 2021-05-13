package utils

import (
	"blog/config"
	"blog/log"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Info struct {
	ID     int64  `json:"id"`
	CardID string `json:"card_id"`
	Expire bool   `json:"expire"`
}

type PayLoad struct {
	jwt.StandardClaims
	Info
}

func GenerateToken(id int64, cardId string) (tokenString string, err error) {
	var claims = &PayLoad{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
		Info: Info{
			ID:     id,
			CardID: cardId,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Logger.Info(config.Settings.SecretKey)
	tokenString, err = token.SignedString([]byte(config.Settings.SecretKey))
	return
}

func ParseToken(tokenString string) (info *Info, err error) {

	token, err := jwt.ParseWithClaims(tokenString, &PayLoad{}, func(*jwt.Token) (interface{}, error) {
		return []byte(config.Settings.SecretKey), nil
	})
	if err != nil {
		log.Logger.Info(err.Error())
		return nil, err
	}
	claims := token.Claims.(*PayLoad)

	var expire bool
	if time.Now().Unix()-claims.ExpiresAt >= 0 {
		expire = true
	}
	info = &Info{
		ID:     claims.ID,
		CardID: claims.CardID,
		Expire: expire,
	}
	return info, err
}
