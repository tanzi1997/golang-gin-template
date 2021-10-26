package dao

import (
	"fmt"
	"golang-gin-template/model"
)

type GatewayServiceInfoDao struct {
}

func NewGatewayServiceInfoDao() *GatewayServiceInfoDao {
	return &GatewayServiceInfoDao{}
}

func (that *GatewayServiceInfoDao) FindList(where *model.GatewayServiceInfo, Offset, Limit int) (int64, []model.GatewayServiceInfo) {
	var count int64
	var list []model.GatewayServiceInfo
	db.Find(where).Count(&count)
	fmt.Println(Offset, Limit)
	db.Limit(Limit).Offset(Offset).Find(&list)

	return count, list
}

func (that *GatewayServiceInfoDao) Create(value *model.GatewayServiceInfo) {
	db.Create(value)
}

func (that *GatewayServiceInfoDao) Update(value *model.GatewayServiceInfo) {
	db.Save(value)
}

func (that *GatewayServiceInfoDao) Delete(id int) {
	var model *model.GatewayServiceInfo
	db.Select(
		"GatewayServiceHttpRule",
		"GatewayServiceTcpRule",
		"GatewayServiceGrpcRule",
		"GatewayServiceAccessControl",
		"GatewayServiceLoadBalance",
	).Unscoped().Delete(&model, id)
}
