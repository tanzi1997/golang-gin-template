package router

import (
	"golang-gin-template/buss"
	"net/http"

	"github.com/gin-gonic/gin"
)

type editServiceUri struct {
	Id int `uri:"id" validate:"required"`
}

type httpServiceUpdateBody struct {
	ServiceName            string `json:"serviceName"`
	ServiceDesc            string `json:"serviceDesc"`
	RuleType               int8   `json:"ruleType"`
	Rule                   string `json:"rule"`
	NeedHttps              int8   `json:"needHttps"`
	NeedStripUri           int8   `json:"needStripUri"`
	NeedWebsocket          int8   `json:"needWebsocket"`
	UrlRewrite             string `json:"urlRewrite"`
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

func editServiceHandler(c *gin.Context) {
	var (
		parmas editServiceUri
		body   httpServiceUpdateBody
	)
	c.ShouldBindUri(&parmas)
	c.ShouldBindJSON(&body)

	validate.Struct(&parmas)
	validate.Struct(&body)

	serviceBuss := buss.NewServiceBuss()

	serviceBuss.Delete(parmas.Id)

	c.JSON(http.StatusOK, gin.H{})
}
