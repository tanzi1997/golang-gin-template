package model

type GatewayServiceGrpcRule struct {
	PrimarykeyIdModel
	ServiceId      int    `gorm:"comment:服务id"`
	Port           int16  `gorm:"comment:端口"`
	HeaderTransfor string `gorm:"type:varchar(5000);comment:header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔"`
	TimeMode
}
