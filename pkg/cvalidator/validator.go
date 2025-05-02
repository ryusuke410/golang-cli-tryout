package cvalidator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func RegisterValidation(name string, fn validator.Func) {
	validate.RegisterValidation(name, fn)
}

func Struct(obj any) error {
	return validate.Struct(obj)
}
