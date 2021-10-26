package buss

import (
	"golang-gin-template/common"
	"golang-gin-template/dao"
	"golang-gin-template/model"
)

type AuthBuss struct {
}

func NewAuthBuss() *AuthBuss {
	that := &AuthBuss{}
	return that
}

func (that *AuthBuss) Register(username, password string) *model.GatewayAdmin {
	GatewayAdminDao := dao.NewGatewayAdminDao()

	admin := &model.GatewayAdmin{
		UserName: username,
		Password: common.HashPassword(password),
	}

	GatewayAdminDao.Create(admin)

	return admin
}

func (that *AuthBuss) Login(username, password string) string {
	GatewayAdminDao := dao.NewGatewayAdminDao()

	admin := &model.GatewayAdmin{
		UserName: username,
	}

	GatewayAdminDao.Find(admin)

	right := common.CheckPasswordHash(password, admin.Password)

	if right {
		panic("错误")
	}

	token, err := common.CreateToken(admin.UserName)

	if err != nil {
		panic("错误")
	}

	return token
}
