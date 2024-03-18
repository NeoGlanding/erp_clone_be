package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemPriceType string

const (
	SALES    ItemPriceType = "SALES"
	PURCHASE ItemPriceType = "PURCHASE"
)

type ItemPrice struct {
	ID string `json:"id" gorm:"primaryKey"`

	ItemId string `json:"item_id"`
	Item   Item   `json:"item:"`

	CurrencyId    string        `json:"currency_id"`
	Currency      Currency      `json:"currency"`
	Price         int64         `json:"price"`
	ItemPriceType ItemPriceType `json:"item_price_type"`

	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}

func (i *ItemPrice) BeforeCreate(tx *gorm.DB) {
	i.ID = uuid.New().String()
}
