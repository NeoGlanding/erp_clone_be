package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Party struct {
	// gorm.Model
	ID				string		`gorm:"primaryKey" json:"id"`
	Name			string		`validate:"required" json:"name"`
	AddressLine1	string		`validate:"required" json:"address_line_1"`
	AddressLine2	*string		`json:"address_line_2"`
	AddressLine3	*string		`json:"address_line_3"`
	PostalCode		string		`validate:"required" json:"postal_code"`
	CountryId		string		`json:"country_id"`
	Country			Country		`json:"country"`
	Users			[]User		`json:"users" gorm:"many2many:user_party_permissions"`
	FileId			*string		`json:"file_id"`
	File			File		`json:"file"`

	CreatedAt		time.Time	`json:"created_at" gorm:"<-:create"`
	UpdatedAt		*time.Time	`json:"updated_at"`
	DeletedAt		*time.Time	`json:"-"`
}

func (p *Party) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}

func (p *Party) AfterSave(tx *gorm.DB) (err error) {
	return err;
}