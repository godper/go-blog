package main

import (
	"blog/conf"
	"blog/http/api"
	"blog/router"
	"blog/service"

	"github.com/gin-gonic/gin"
)

func main() {
	//配置初始化
	conf.Init()
	//服务初始化
	service.Init()
	//api初始化
	api.Init()

	//运行模式
	gin.SetMode(conf.Cfg.Server.RunMode)

	e := gin.Default()
	//路由初始化
	router.Init(e)

	e.Run(":3000")
}
