package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerPartnership struct {
	ID string `gorm:"primaryKey" json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`json:"-"`
}

func (c *CustomerPartnership) BeforeCreate() (error) {
	c.ID = uuid.New().String()
	return nil
}