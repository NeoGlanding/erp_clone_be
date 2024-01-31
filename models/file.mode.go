package models

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	Id					string 		`json:"id" gorm:"primaryKey"`
	Filename			string		`json:"-"`
	UserId				string		`gorm:"unique" json:"user_id"`
	User				User		`json:"-"`
	FileUrl				string		`json:"url"`

	CreatedAt	time.Time			`json:"created_at" gorm:"<-:create"`
	UpdatedAt	*time.Time			`json:"updated_at"`
	DeletedAt	gorm.DeletedAt		`json:"-"`
}

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	f.Id = uuid.New().String()

	
	return
}

func (f *File) AfterCreate(tx *gorm.DB) (err error) {
	bucketName := os.Getenv("FIREBASE_BUCKET_URL")

	f.FileUrl = fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", bucketName, f.Filename, f.Filename)

	return nil
}