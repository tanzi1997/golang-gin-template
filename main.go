package main

import (
	"golang-gin-template/common"
	"golang-gin-template/dao"
	"golang-gin-template/middleware"
	"golang-gin-template/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 配置文件初始化
	common.ConfigInit()
	// 数据库初始化
	dao.Init(common.Config.Mysql)
	// 初始化
	c := gin.Default()
	// cors
	c.Use(middleware.Cors())
	// 注册router
	router.InitRouter(c)
	// 监听端口
	s := &http.Server{
		Addr:           ":8080",
		Handler:        c,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
