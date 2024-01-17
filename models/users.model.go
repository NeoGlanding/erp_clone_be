package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type User struct {
	Id			string 		`gorm:"primaryKey"`
	Email		string			
	Password	string
	Phone		*string
	CreatedAt	time.Time
	UpdatedAt	*time.Time
	DeletedAt	*time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	return
}