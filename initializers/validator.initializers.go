package initializers

import (
	"github.com/automa8e_clone/libs"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func Validator() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
	Validate.RegisterValidation("datestring", libs.DateValidator)
}