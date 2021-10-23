package router

import (
	"net/http"
	"tanzi1997/buss"

	"github.com/gin-gonic/gin"
)

type createGrpcServiceBody struct {
	ServiceName            string `json:"serviceName"`
	ServiceDesc            string `json:"serviceDesc"`
	Port                   int16  `json:"port"`
	HeaderTransfor         string `json:"headerTransfor"`
	OpenAuth               int8   `json:"openAuth"`
	WhiteList              string `json:"whiteList"`
	ClientipFlowLimit      int32  `json:"clientipFlowLimit"`
	ServiceFlowLimit       int32  `json:"serviceFlowLimit"`
	RoundType              int8   `json:"roundType"`
	IpList                 string `json:"ipList"`
	WeightList             string `json:"weightList"`
	UpstreamConnectTimeout int32  `json:"upstreamConnectTimeout"`
	UpstreamHeaderTimeout  int32  `json:"upstreamHeaderTimeout"`
	UpstreamIdleTimeout    int32  `json:"upstreamIdleTimeout"`
	UpstreamMaxIdle        int32  `json:"upstreamMaxIdle"`
}

func createGrpcerviceHandler(c *gin.Context) {
	var body createGrpcServiceBody
	c.ShouldBindJSON(&body)
	validate.Struct(&body)

	serviceBuss := buss.NewGrpcServiceBuss()

	serviceInfo := serviceBuss.CreateRpc(
		body.ServiceName,
		body.ServiceDesc,
		body.Port,
		body.HeaderTransfor,
		body.RoundType,
		body.IpList,
		body.WeightList,
		body.UpstreamConnectTimeout,
		body.UpstreamHeaderTimeout,
		body.UpstreamIdleTimeout,
		body.UpstreamMaxIdle,
	)

	c.JSON(http.StatusOK, serviceInfo)
}
