package dao

import (
	"tanzi1997/model"
)

type GatewayServiceTcpRuleDao struct {
}

func NewGatewayServiceTcpRuleDao() *GatewayServiceTcpRuleDao {
	return &GatewayServiceTcpRuleDao{}
}

func (that *GatewayServiceTcpRuleDao) Create(value *model.GatewayServiceTcpRule) {
	db.Create(value)
}

func (that *GatewayServiceTcpRuleDao) FindList(where *model.GatewayServiceTcpRule, count *int64, Limit, Offset int) {
	db.Find(where).Count(count)
	db.Limit(Limit).Offset(Offset).Find(where)
}
