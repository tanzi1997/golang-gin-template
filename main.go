package main

import (
	"net/http"
	"tanzi1997/common"
	"tanzi1997/dao"
	"tanzi1997/router"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 可将将* 替换为指定的域名
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func main() {
	// 配置文件初始化
	common.ConfigInit()
	// 数据库初始化
	dao.Init(common.Config.Mysql)
	// 初始化
	c := gin.Default()
	// cors
	c.Use(Cors())
	// 注册router
	router.InitRouter(c)
	// 监听端口
	c.Run(":8180")
}
