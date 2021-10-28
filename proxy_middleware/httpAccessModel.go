package proxy_middleware

import "github.com/gin-gonic/gin"

type ServiceManager struct {
	ServiceMap map[string]
}

// 匹配接入方式 基于请求信息
func HttpAccessModelMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
