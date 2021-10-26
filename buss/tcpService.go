package buss

import (
	"golang-gin-template/dao"
	"golang-gin-template/model"
)

type TcpServiceBuss struct {
}

func NewTcpServiceBuss() *TcpServiceBuss {
	return &TcpServiceBuss{}
}

func (that *TcpServiceBuss) CreateTcp(
	serviceDesc,
	serviceName string,
	Port int16,
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
		GatewayServiceTcpRule: model.GatewayServiceTcpRule{
			Port: Port,
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
