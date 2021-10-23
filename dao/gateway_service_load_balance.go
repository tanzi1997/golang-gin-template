package dao

import (
	"tanzi1997/model"
)

type GatewayServiceLoadBalanceDao struct {
}

func NewGatewayServiceLoadBalanceDao() *GatewayServiceLoadBalanceDao {
	return &GatewayServiceLoadBalanceDao{}
}

func (that *GatewayServiceLoadBalanceDao) Create(value *model.GatewayServiceLoadBalance) {
	db.Create(value)
}

func (that *GatewayServiceLoadBalanceDao) FindList(where *model.GatewayServiceLoadBalance, count *int64, Limit, Offset int) {
	db.Find(where).Count(count)
	db.Limit(Limit).Offset(Offset).Find(where)
}
