package buss

import (
	"golang-gin-template/dao"
	"golang-gin-template/model"
)

type HttpServiceBuss struct {
}

func NewHttpServiceBuss() *HttpServiceBuss {
	return &HttpServiceBuss{}
}

func (that *HttpServiceBuss) CreateHttp(
	serviceDesc,
	serviceName string,
	RuleType int8,
	Rule string,
	NeedHttps,
	NeedStripUri,
	NeedWebsocket int8,
	UrlRewrite string,
	HeaderTransfor string,
	RoundType int8,
	IpList,
	WeightList string,
	UpstreamConnectTimeout,
	UpstreamHeaderTimeout,
	UpstreamIdleTimeout,
	UpstreamMaxIdle int32,
) *model.GatewayServiceInfo {
	gatewayServiceInfoDao := dao.NewGatewayServiceInfoDao()

	serviceInfo := &model.GatewayServiceInfo{
		LoadType:    0,
		ServiceName: serviceName,
		ServiceDesc: serviceDesc,
		GatewayServiceHttpRule: model.GatewayServiceHttpRule{
			RuleType:       RuleType,
			Rule:           Rule,
			NeedHttps:      NeedHttps,
			NeedStripUri:   NeedStripUri,
			NeedWebsocket:  NeedWebsocket,
			UrlRewrite:     UrlRewrite,
			HeaderTransfor: HeaderTransfor,
		},
		GatewayServiceLoadBalance: model.GatewayServiceLoadBalance{
			RoundType:              RoundType,
			IpList:                 IpList,
			WeightList:             WeightList,
			UpstreamConnectTimeout: UpstreamConnectTimeout,
			UpstreamHeaderTimeout:  UpstreamHeaderTimeout,
			UpstreamIdleTimeout:    UpstreamIdleTimeout,
			UpstreamMaxIdle:        UpstreamMaxIdle,
		},
	}

	gatewayServiceInfoDao.Create(serviceInfo)

	return serviceInfo
}
