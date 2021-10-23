package router

import (
	"net/http"
	"tanzi1997/buss"

	"github.com/gin-gonic/gin"
)

type createHttpServiceBody struct {
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

// type GatewayServiceHttpRuleOutPut struct {
// 	Id             int            `json:"id"`
// 	ServiceId      int            `json:"service_id"`
// 	RuleType       int8           `json:"rule_type"`
// 	Rule           string         `json:"rule"`
// 	NeedHttps      int8           `json:"need_https"`
// 	NeedStripUri   int8           `json:"need_strip_uri"`
// 	NeedWebsocket  int8           `json:"need_websocket"`
// 	UrlRewrite     string         `json:"url_rewrite"`
// 	HeaderTransfor string         `json:"header_transfor"`
// 	CreatedAt      time.Time      `json:"created_at"`
// 	UpdatedAt      time.Time      `json:"updated_at"`
// 	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
// }

// type GatewayServiceLoadBalanceOutPut struct {
// 	Id                     int            `json:"id"`
// 	ServiceId              int            `json:"service_id"`
// 	CheckMethod            int            `json:"check_method"`
// 	CheckTimeout           int16          `json:"check_timeout"`
// 	CheckInterval          int32          `json:"check_interval"`
// 	RoundType              int8           `json:"round_type"`
// 	IpList                 string         `json:"ip_list"`
// 	WeightList             string         `json:"weight_list"`
// 	ForbidList             string         `json:"forbid_list"`
// 	UpstreamConnectTimeout int32          `json:"upstream_connect_timeout"`
// 	UpstreamHeaderTimeout  int32          `json:"upstream_header_timeout"`
// 	UpstreamIdleTimeout    int32          `json:"upstream_idle_timeout"`
// 	UpstreamMaxIdle        int32          `json:"upstream_MaxIdle"`
// 	CreatedAt              time.Time      `json:"created_at"`
// 	UpdatedAt              time.Time      `json:"updated_at"`
// 	DeletedAt              gorm.DeletedAt `json:"deleted_at"`
// }

// type httpServiceCreateOutPut struct {
// 	Id                 int                             `json:"id"`
// 	LoadType           int8                            `json:"load_type"`
// 	ServiceName        string                          `json:"service_name"`
// 	ServiceDesc        string                          `json:"service_desc"`
// 	ServiceHttpRule    GatewayServiceHttpRuleOutPut    `json:"service_http_rule"`
// 	ServiceLoadBalance GatewayServiceLoadBalanceOutPut `json:"service_load_balance"`
// 	CreatedAt          time.Time                       `json:"created_at"`
// 	UpdatedAt          time.Time                       `json:"updated_at"`
// 	DeletedAt          gorm.DeletedAt                  `json:"deleted_at"`
// }

func createHttpServiceHandler(c *gin.Context) {
	var body createHttpServiceBody
	c.ShouldBindJSON(&body)
	validate.Struct(&body)

	serviceBuss := buss.NewHttpServiceBuss()

	serviceInfo := serviceBuss.CreateHttp(
		body.ServiceName,
		body.ServiceDesc,
		body.RuleType,
		body.Rule,
		body.NeedHttps,
		body.NeedStripUri,
		body.NeedWebsocket,
		body.UrlRewrite,
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
