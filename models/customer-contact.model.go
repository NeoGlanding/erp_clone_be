package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerContact struct {
	ID string `json:"id" gorm:"primaryKey"`

	CustomerId string   `json:"customer_id"`
	Customer   Customer `json:"-"`

	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (c *CustomerContact) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}
