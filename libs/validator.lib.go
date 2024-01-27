package libs

import (
	"regexp"

	"github.com/automa8e_clone/helpers"
	"github.com/go-playground/validator/v10"
)

func DateValidator(fl validator.FieldLevel) bool {
	valid := helpers.IsValidDateString(fl.Field().String())
	return valid
}

func PasswordValidator(fl validator.FieldLevel) bool {

	criteria1 := "[0-9]+"
	criteria2 := "[a-z]+"
	criteria3 := "[A-Z]+"

	rgxC1, _ := regexp.Compile(criteria1); 
	rgxC2, _ := regexp.Compile(criteria2); 
	rgxC3, _ := regexp.Compile(criteria3); 

	match := rgxC1.Match([]byte(fl.Field().String()));
	match2 := rgxC2.Match([]byte(fl.Field().String()));
	match3 := rgxC3.Match([]byte(fl.Field().String()));

	return match && match2 && match3

}