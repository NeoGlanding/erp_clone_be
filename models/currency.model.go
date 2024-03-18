package models

import (
	"time"

	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

type Currency struct {
	ID string `json:"id" gorm:"primaryKey"`

	Currency string `json:"currency"`

	CreatedAt time.Time       `json:"created_at" gorm:"<-:create"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}

func (c *Currency) TableName() string {
	return "currencies"
}

type CurrencySeeder struct {
	gorm_seeder.SeederAbstract
}

func NewCurrencySeeder(cfg gorm_seeder.SeederConfiguration) CurrencySeeder {
	return CurrencySeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

func (c *CurrencySeeder) Seed(db *gorm.DB) error {
	currencies := []Currency{
		{ID: "4c062baa-96c8-4160-9b98-f91bdcbfb655", Currency: "SGD"},
		{ID: "2b42904f-b039-41f1-80ec-788e500ec9b3", Currency: "MYR"},
		{ID: "5213b5ff-964b-4a1f-a5e8-f9b6a6a6eff5", Currency: "IDR"},
		// Add more currencies as needed
	}

	for _, currency := range currencies {
		db.FirstOrCreate(&currency, currency)
	}

	return nil
}

func (c *CurrencySeeder) Clear(db *gorm.DB) error {
	entity := Currency{}
	return c.SeederAbstract.Delete(db, entity.TableName())
}
