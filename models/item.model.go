package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID string `json:"id" gorm:"primaryKey"`

	CreatedByUser   User   `json:"-"`
	CreatedByUserId string `json:"-"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`

	PartyId string `json:"party_id"`
	Party   Party  `json:"party"`

	Active bool `json:"active"`

	ItemPrice []ItemPrice `json:"item_prices" gorm:"foreignKey:ItemId"`

	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) {
	i.ID = uuid.New().String()
}
