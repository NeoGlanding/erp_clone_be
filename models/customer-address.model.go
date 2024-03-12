package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerAddresses struct {
	ID string `json:"id" gorm:"primary_key"`

	AddressLine1 string  `json:"address_line1"`
	AddressLine2 string  `json:"address_line2"`
	AddressLine3 string  `json:"address_line3"`
	PostalCode   string  `json:"postal_code"`
	Country      Country `json:"country"`

	Customer   Customer `json:"-"`
	CustomerId string   `json:"customer_id"`
	CountryId  string   `json:"country_id"`

	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}

func (c *CustomerAddresses) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
