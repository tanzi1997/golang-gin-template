package dao

import (
	"tanzi1997/model"
)

type GatewayAdminDao struct {
}

func NewGatewayAdminDao() *GatewayAdminDao {
	return &GatewayAdminDao{}
}

func (that *GatewayAdminDao) Create(a *model.GatewayAdmin) {
	db.Create(a)
}

func (that *GatewayAdminDao) Find(where *model.GatewayAdmin) {
	db.Where(where)
}
