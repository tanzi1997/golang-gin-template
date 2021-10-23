package router

import (
	"net/http"
	"tanzi1997/buss"

	"github.com/gin-gonic/gin"
)

type deleteService struct {
	Id int `uri:"id" validate:"required"`
}

func deleteServiceHandler(c *gin.Context) {
	var parmas deleteService
	c.ShouldBindUri(&parmas)
	validate.Struct(&parmas)

	serviceBuss := buss.NewServiceBuss()

	serviceBuss.Delete(parmas.Id)

	c.JSON(http.StatusOK, gin.H{})
}
