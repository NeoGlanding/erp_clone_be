package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type UserDetails struct {
	Id					string 		`gorm:"primaryKey"`
	UserId				string		`gorm:"unique"`
	User				User			
	FirstName			string
	Surname				*string
	AddressLine1		string
	AddressLine2		string
	AddressLine3		string
	PostalCode			string
	CountryId			string
	Country				Country
	DateOfBirth			time.Time


	CreatedAt	time.Time
	UpdatedAt	*time.Time
	DeletedAt	gorm.DeletedAt
}

func (u *UserDetails) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	return
}