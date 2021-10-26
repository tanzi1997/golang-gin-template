package router

import (
	"golang-gin-template/buss"
	"net/http"

	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func loginHandler(c *gin.Context) {
	var obj login
	c.ShouldBindJSON(&obj)
	validate.Struct(&obj)

	authBuss := buss.NewAuthBuss()

	tokeng := authBuss.Login(obj.Username, obj.Password)

	c.JSON(http.StatusOK, gin.H{"tokeng": tokeng})
}
