package http_proxyr_outer

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(c *gin.Engine) {
	// 初始化
	c.GET("/ping")

}
