package model

type GatewayServiceLoadBalance struct {
	PrimarykeyIdModel
	ServiceId              int    `gorm:"comment:服务id"`
	CheckMethod            int    `gorm:"comment:检查方法 0=tcpchk,检测端口是否握手成功"`
	CheckTimeout           int16  `gorm:"comment:check超时时间,单位s"`
	CheckInterval          int32  `gorm:"comment:检查间隔, 单位s"`
	RoundType              int8   `gorm:"comment:轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash"`
	IpList                 string `gorm:"type:varchar(2000);comment:ip列表"`
	WeightList             string `gorm:"type:varchar(2000);comment:权重列表"`
	ForbidList             string `gorm:"type:varchar(2000);comment:禁用ip列表"`
	UpstreamConnectTimeout int32  `gorm:"comment:建立连接超时, 单位s"`
	UpstreamHeaderTimeout  int32  `gorm:"comment:获取header超时, 单位s"`
	UpstreamIdleTimeout    int32  `gorm:"comment:链接最大空闲时间, 单位s"`
	UpstreamMaxIdle        int32  `gorm:"comment:最大空闲链接数"`
	TimeMode
}
