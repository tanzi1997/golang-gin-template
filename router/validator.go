package router

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidate() {
	validate = validator.New()
}
