package router

import (
	"golang-gin-template/buss"
	"net/http"

	"github.com/gin-gonic/gin"
)

type register struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func registerHandler(c *gin.Context) {
	var obj register
	c.ShouldBindJSON(&obj)
	validate.Struct(&obj)

	authBuss := buss.NewAuthBuss()

	createDate := authBuss.Register(obj.Username, obj.Password)

	c.JSON(http.StatusCreated, createDate)
}
