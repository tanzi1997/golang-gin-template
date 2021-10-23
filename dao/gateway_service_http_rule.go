package dao

import (
	"tanzi1997/model"
)

type GatewayServiceHttpRuleDao struct {
}

func NewGatewayServiceHttpRuleDao() *GatewayServiceHttpRuleDao {
	return &GatewayServiceHttpRuleDao{}
}

func (that *GatewayServiceHttpRuleDao) Create(value *model.GatewayServiceHttpRule) {
	db.Create(value)
}

func (that *GatewayServiceHttpRuleDao) FindList(where *model.GatewayServiceHttpRule, count *int64, Limit, Offset int) {
	db.Find(where).Count(count)
	db.Limit(Limit).Offset(Offset).Find(where)
}
