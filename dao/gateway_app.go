package dao

import (
	"golang-gin-template/model"
)

type GatewayAppDao struct {
}

func NewGatewayAppDao() *GatewayAppDao {
	return &GatewayAppDao{}
}

func (that *GatewayAppDao) Create(value *model.GatewayApp) {
	db.Create(value)
}

func (that *GatewayAppDao) FindList(where *model.GatewayApp, count *int64, Limit, Offset int) {
	db.Find(where).Count(count)
	db.Limit(Limit).Offset(Offset).Find(where)
}
