package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id                 string  `json:"id" gorm:"primaryKey"`
	Email              string  `json:"email" gorm:"unique"`
	Password           string  `json:"-"`
	Phone              *string `json:"phone" gorm:"unique"`
	Party              []Party `json:"-" gorm:"many2many:user_party_permissions"`
	InformationChanged int     `gorm:"default:0"`

	Customer []Customer `json:"-" gorm:"foreignKey:CreatedByUserId"`
	Item     []Item     `json:"-" gorm:"foreignKey:CreatedByUserId"`

	CreatedAt time.Time  `json:"-" gorm:"<-:create"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	return
}
