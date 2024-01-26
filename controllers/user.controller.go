package controllers

import (
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PostOnboardingBody struct {
	Email			string		`json:"email" validate:"required,email"`
	FirstName		string		`json:"first_name" validate:"required,max=100"`
	Surname			string		`json:"surname" validate:"required,max=100"`
	AddressLine1	string		`json:"address_line_1" validate:"required,max=100"`
	AddressLine2	*string		`json:"address_line_2" validate:"omitempty,max=100"`
	AddressLine3	*string		`json:"address_line_3" validate:"omitempty,max=100"`
	PostalCode		string		`json:"postal_code" validate:"required,max=100"`
	CountryId		string		`json:"country_id" validate:"required"`
	IdentityNumber	string		`json:"identity_number" validate:"required"`
	DateOfBirth		string		`json:"date_of_birth" validate:"required,datestring"`
}

func OnboardUser(c *gin.Context) {
	var body PostOnboardingBody;

	c.ShouldBindBodyWith(&body, binding.JSON)

	err := initializers.Validate.Struct(body);

	if err != nil {
		helpers.SetValidationError(c, &err)
		return
	}

	// var user models.User

	// exist := users_repository.FindByEmail(body.Email, &user)

	

}