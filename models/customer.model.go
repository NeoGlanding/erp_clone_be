package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID string `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	BusinessRegistrationNumber *string `json:"business_registration_number"`
	Url *string `json:"url"`
	Remarks *string `json:"remarks"`

	CustomerTypeId string `json:"customer_type_id"`
	CustomerPartnershipId string `json:"customer_partnership_id"`
	CountryId string `json:"country_id"`
	PartyId string `json:"party_id"`

	FileId string `json:"file_id"`

	CreatedByUserId string `json:"created_by_user_id"`

	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`json:"-"`
}

func (*Customer) TableName() string {
	return "customers"
}

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}