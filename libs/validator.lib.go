package libs

import (
	"github.com/automa8e_clone/helpers"
	"github.com/go-playground/validator/v10"
)

func DateValidator(fl validator.FieldLevel) bool {
	valid := helpers.IsValidDateString(fl.Field().String())
	if valid {
		return true
	} else {
		return false
	}
	
}