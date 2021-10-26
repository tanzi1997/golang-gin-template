package buss

import (
	"golang-gin-template/dao"
	"golang-gin-template/model"
)

type ServiceBuss struct {
}

func NewServiceBuss() *ServiceBuss {
	that := &ServiceBuss{}
	return that
}

func (that *ServiceBuss) GetList(query string, offset, limit int) (int64, []model.GatewayServiceInfo) {
	gatewayServiceInfoDao := dao.NewGatewayServiceInfoDao()

	serviceInfo := &model.GatewayServiceInfo{
		// ServiceName: query,
		// ServiceDesc: query,
	}

	count, gatewayServiceInfoList := gatewayServiceInfoDao.FindList(serviceInfo, offset, limit)

	return count, gatewayServiceInfoList
}

func (that *ServiceBuss) Delete(id int) {
	gatewayServiceInfoDao := dao.NewGatewayServiceInfoDao()
	gatewayServiceInfoDao.Delete(id)
}

func (that *ServiceBuss) CreateGrcp(serviceDesc, serviceName string) {
	// gatewayServiceInfoDao := dao.NewGatewayServiceInfoDao()
	// GatewayServiceGrpcRuleDao := dao.NewGatewayServiceGrpcRuleDao()
	// gatewayServiceLoadBalanceDao := dao.NewGatewayServiceLoadBalanceDao()

	// serviceInfo := &model.GatewayServiceInfo{
	// 	LoadType:    0,
	// 	ServiceName: serviceName,
	// 	ServiceDesc: serviceDesc,
	// }
}

func (that *ServiceBuss) UpdateHttp(
	id int,
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
