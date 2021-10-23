package dao

import (
	"tanzi1997/model"
)

type GatewayServiceGrpcRuleDao struct {
}

func NewGatewayServiceGrpcRuleDao() *GatewayServiceGrpcRuleDao {
	return &GatewayServiceGrpcRuleDao{}
}

func (that *GatewayServiceGrpcRuleDao) Create(value *model.GatewayServiceGrpcRule) {
	db.Create(value)
}

func (that *GatewayServiceGrpcRuleDao) FindList(where *model.GatewayServiceGrpcRule, count *int64, Limit, Offset int) {
	db.Find(where).Count(count)
	db.Limit(Limit).Offset(Offset).Find(where)
}
