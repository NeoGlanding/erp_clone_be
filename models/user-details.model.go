package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type UserDetails struct {
	Id			string 		`gorm:"primaryKey"`
	UserId		string
	User		User			
	FirstName	string
	Surname		*string
	Address		*string
	CreatedAt	time.Time
	UpdatedAt	*time.Time
	DeletedAt	*time.Time
}

func (u *UserDetails) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	return
}