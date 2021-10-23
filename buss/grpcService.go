package buss

import (
	"tanzi1997/dao"
	"tanzi1997/model"
)

type GrpcServiceBuss struct {
}

func NewGrpcServiceBuss() *GrpcServiceBuss {
	return &GrpcServiceBuss{}
}

func (that *GrpcServiceBuss) CreateRpc(
	serviceDesc,
	serviceName string,
	Port int16,
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
		GatewayServiceGrpcRule: model.GatewayServiceGrpcRule{
			Port:           Port,
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
