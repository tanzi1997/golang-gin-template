package model

type GatewayServiceInfo struct {
	PrimarykeyIdModel
	LoadType                    int8                        `gorm:"comment:负载类型 0=http 1=tcp 2=grpc"`
	ServiceName                 string                      `gorm:"type:varchar(255);comment:服务名称 6-128 数字字母下划线"`
	ServiceDesc                 string                      `gorm:"type:varchar(255);comment:服务描述"`
	GatewayServiceHttpRule      GatewayServiceHttpRule      `gorm:"foreignKey:ServiceId"`
	GatewayServiceTcpRule       GatewayServiceTcpRule       `gorm:"foreignKey:ServiceId"`
	GatewayServiceGrpcRule      GatewayServiceGrpcRule      `gorm:"foreignKey:ServiceId"`
	GatewayServiceAccessControl GatewayServiceAccessControl `gorm:"foreignKey:ServiceId"`
	GatewayServiceLoadBalance   GatewayServiceLoadBalance   `gorm:"foreignKey:ServiceId"`
	TimeMode
}
