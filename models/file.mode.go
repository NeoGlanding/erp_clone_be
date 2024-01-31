package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	Id					string 		`json:"id" gorm:"primaryKey"`
	UserId				string		`gorm:"unique"`
	User				User		`json:"credentials"`	

	CreatedAt	time.Time			`json:"created_at" gorm:"<-:create"`
	UpdatedAt	*time.Time			`json:"updated_at"`
	DeletedAt	gorm.DeletedAt		`json:"-"`
}

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	f.Id = uuid.New().String()
	return
}