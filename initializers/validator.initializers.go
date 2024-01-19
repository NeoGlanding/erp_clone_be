package initializers

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func Validator() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}