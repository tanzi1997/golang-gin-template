package model

type GatewayAdmin struct {
	PrimarykeyIdModel
	UserName string `gorm:"type:varchar(255);comment:用户名"`
	Password string `gorm:"type:varchar(255);comment:密码"`
	TimeMode
}
