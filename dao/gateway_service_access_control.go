package dao

import (
	"tanzi1997/model"
)

type GatewayServiceAccessControlDao struct {
}

func NewGatewayServiceAccessControlDao() *GatewayServiceAccessControlDao {
	return &GatewayServiceAccessControlDao{}
}

func (that *GatewayServiceAccessControlDao) Create(value *model.GatewayServiceAccessControl) {
	db.Create(value)
}

func (that *GatewayServiceAccessControlDao) FindList(where *model.GatewayServiceAccessControl, count *int64, Limit, Offset int) {
	db.Find(where).Count(count)
	db.Limit(Limit).Offset(Offset).Find(where)
}
