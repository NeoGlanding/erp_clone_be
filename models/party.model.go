package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Party struct {
	gorm.Model
	ID				string		`gorm:"primaryKey" json:"id"`
	Name			string		`validate:"required" json:"name"`
	AddressLine1	string		`validate:"required" json:"address_line_1"`
	AddressLine2	*string		`json:"address_line_2"`
	AddressLine3	*string		`json:"address_line_3"`
	PostalCode		string		`validate:"required" json:"postal_code"`
	CountryId		string		`json:"country_id"`
	Country			Country

	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		*time.Time	`json:"updated_at"`
	DeletedAt		*time.Time	
}

func (p *Party) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	return
}