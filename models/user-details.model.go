package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type UserDetails struct {
	Id								string 		`json:"id" gorm:"primaryKey"`
	UserId							string		`gorm:"unique"`
	User							User		`json:"-"`	
	FirstName						string		`json:"first_name"`
	Surname							string		`json:"surname"`
	AddressLine1					string		`json:"address_line_1"`
	AddressLine2					string		`json:"address_line_2"`
	AddressLine3					string		`json:"address_line_3"`
	PostalCode						string		`json:"postal_code"`
	CountryId						string		`json:"country_id"`
	Country							Country		`json:"country"`
	IdentityNumber					string		`json:"identity_number"`
	DateOfBirth						time.Time	`json:"date_of_birth"`
	ProfilePictureFileId			*string		`json:"-"`
	ProfilePictureFile				File		`json:"-"`


	CreatedAt	time.Time			`json:"created_at" gorm:"<-:create"`
	UpdatedAt	*time.Time			`json:"updated_at"`
	DeletedAt	gorm.DeletedAt		`json:"-"`
}

func (u *UserDetails) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	return
}