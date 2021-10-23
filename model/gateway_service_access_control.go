package model

type GatewayServiceAccessControl struct {
	PrimarykeyIdModel
	ServiceId         int    `gorm:"comment:服务id"`
	OpenAuth          int8   `gorm:"comment:是否开启权限 1=开启"`
	BlackList         string `gorm:"type:varchar(1000);comment:黑名单ip"`
	WhiteList         string `gorm:"type:varchar(1000);comment:白名单ip"`
	WhiteHostName     string `gorm:"type:varchar(1000);comment:白名单主机"`
	ClientipFlowLimit int32  `gorm:"comment:客户端ip限流"`
	ServiceFlowLimit  int32  `gorm:"comment:服务端限流"`
	TimeMode
}
