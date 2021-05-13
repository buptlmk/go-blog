package main

import (
	"blog/config"
	"blog/database/mysql"
	"blog/database/redis"
	"blog/log"
	"blog/router"
	_ "blog/router/article"
	_ "blog/router/chat"
	_ "blog/router/comment"
	_ "blog/router/ticket"
	_ "blog/router/user"
	"net/http"
	"strconv"
)

func main() {

	// 加载配置
	config.Load("config.json")

	log.StartLog()

	//
	if err := mysql.InitDB(); err != nil {
		panic(err)
	}
	log.Logger.Info("mysql init success...")
	if err := redis.InitRedis(); err != nil {
		panic(err)
	}
	log.Logger.Info("redis init success...")

	server := &http.Server{
		Addr:    config.Settings.Server.Addr + ":" + strconv.Itoa(config.Settings.Server.Port),
		Handler: router.App,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
