package model

type GatewayServiceTcpRule struct {
	PrimarykeyIdModel
	ServiceId int   `gorm:"comment:服务id"`
	Port      int16 `gorm:"comment:端口号"`
	TimeMode
}
