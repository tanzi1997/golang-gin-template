package dao

import (
	"fmt"
	"golang-gin-template/common"
	"golang-gin-template/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init(config common.Mysql) {
	DSN := config.Username +
		":" +
		config.Password +
		"@tcp(" +
		config.Host +
		":" +
		config.Port + ")/" +
		config.Dbname +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	fmt.Println(DSN)

	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.GatewayAdmin{})
	db.AutoMigrate(&model.GatewayApp{})
	db.AutoMigrate(&model.GatewayServiceAccessControl{})
	db.AutoMigrate(&model.GatewayServiceGrpcRule{})
	db.AutoMigrate(&model.GatewayServiceHttpRule{})
	db.AutoMigrate(&model.GatewayServiceInfo{})
	db.AutoMigrate(&model.GatewayServiceLoadBalance{})
	db.AutoMigrate(&model.GatewayServiceTcpRule{})
}
