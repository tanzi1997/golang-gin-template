package model

type GatewayApp struct {
	PrimarykeyIdModel
	AppId     string `gorm:"type:varchar(255);comment:租户id"`
	Name      string `gorm:"type:varchar(255);comment:租户名称"`
	Secret    string `gorm:"type:varchar(255);comment:密钥"`
	White_ips string `gorm:"type:varchar(1000);comment:ip白名单，支持前缀匹配"`
	Qpd       int    `gorm:"comment:日请求量限制"`
	Qps       int    `gorm:"omment:每秒请求量限制"`
	TimeMode
}
