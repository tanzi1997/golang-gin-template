package router

import (
	"net/http"
	"tanzi1997/buss"

	"github.com/gin-gonic/gin"
)

type serviceDetail struct {
	Query  string `json:"query" validate:"required"`
	Limit  int    `json:"limit" validate:"required"`
	Offset int    `json:"offset" validate:"required"`
}

func serviceDetailHandler(c *gin.Context) {
	var query serviceDetail
	c.ShouldBindQuery(&query)
	validate.Struct(&query)

	serviceBuss := buss.NewServiceBuss()

	list, count := serviceBuss.GetList(query.Query, query.Limit, query.Offset)

	c.JSON(http.StatusOK, gin.H{"list": list, "count": count})
}
