package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(c *gin.Engine) {
	// 初始化
	InitValidate()
	v1 := c.Group("/v1")
	// auth
	v1.POST("/register", registerHandler) // 注册
	v1.POST("/login", loginHandler)       // 登录
	v1.POST("/refresh", loginHandler)     // 刷新token
	// servace
	v1.GET("/service", serviceListHandler)          // 获取
	v1.GET("/service/:id", serviceDetailHandler)    // 获取详情
	v1.PUT("/service/:id", editServiceHandler)      // 修改
	v1.DELETE("/service/:id", deleteServiceHandler) // 删除
	// 添加HTTP服务接口
	v1.POST("/httpService", createHttpServiceHandler) // 创建
	v1.POST("/tcpService", createTcpServiceHandler)   // 创建
	v1.POST("/grcpService", createGrpcerviceHandler)  // 创建
	// 升级HTTP服务接口
	// v1.PUT("/httpService/:id", updatehttpServiceHandler) // 创建
	// v1.PUT("/tcpService/:id", updateTcpServiceHandler)   // 创建
	// v1.PUT("/grcpService/:id", updateGrpcServiceHandler) // 创建
}
