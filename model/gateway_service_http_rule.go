package model

type GatewayServiceHttpRule struct {
	PrimarykeyIdModel
	ServiceId      int    `gorm:"comment:服务id"`
	RuleType       int8   `gorm:"comment:匹配类型 0=url前缀url_prefix 1=域名domain"`
	Rule           string `gorm:"type:varchar(255);comment:type=domain表示域名，type=url_prefix时表示url前缀"`
	NeedHttps      int8   `gorm:"comment:支持https 1=支持"`
	NeedStripUri   int8   `gorm:"comment:启用strip_uri 1=启用"`
	NeedWebsocket  int8   `gorm:"comment:是否支持websocket 1=支持"`
	UrlRewrite     string `gorm:"type:varchar(5000);comment:url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔	"`
	HeaderTransfor string `gorm:"type:varchar(5000);comment:header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔"`
	TimeMode
}
