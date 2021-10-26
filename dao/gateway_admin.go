package dao

import (
	"golang-gin-template/model"
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
