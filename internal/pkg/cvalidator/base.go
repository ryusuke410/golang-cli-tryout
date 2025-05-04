package cvalidator

import (
	"github.com/go-playground/validator/v10"
)

func baseValidation(fl validator.FieldLevel) bool {
	base := fl.Field().Int()
	return base == 0 || (2 <= base && base <= 36)
}

func init() {
	RegisterValidation("c_base", baseValidation)
}
