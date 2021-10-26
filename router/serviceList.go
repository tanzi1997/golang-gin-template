package router

import (
	"golang-gin-template/buss"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serviceListInput struct {
	Query  string `form:"query" validate:"required"`
	Limit  int    `form:"limit" validate:"required"`
	Offset int    `form:"offset" validate:"required"`
}

// type gatewayService struct {
// 	ID          int            `json:"id"`
// 	LoadType    int8           `json:"load_type"`
// 	ServiceName string         `json:"Service_name"`
// 	ServiceDesc string         `json:"Service_desc"`
// 	CreatedAt   time.Time      `json:"created_at"`
// 	UpdatedAt   time.Time      `json:"updated_at"`
// 	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
// }

// type serviceListOutput struct {
// 	Count int            `json:"count"`
// 	List  gatewayService `json:"list"`
// }

func serviceListHandler(c *gin.Context) {
	var query serviceListInput
	c.ShouldBindQuery(&query)
	validate.Struct(&query)

	serviceBuss := buss.NewServiceBuss()

	count, list := serviceBuss.GetList(query.Query, query.Offset, query.Limit)

	c.JSON(http.StatusOK, gin.H{"list": list, "count": count})
}
