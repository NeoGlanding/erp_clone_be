package models

import (
	"time"

	"github.com/google/uuid"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

type CustomerPartnership struct {
	ID string `gorm:"primaryKey" json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	Customer []Customer `gorm:"foreignKey:CustomerPartnershipId" json:"customer"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (c *CustomerPartnership) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

func (c *CustomerPartnership) TableName() string {
	return "customer_partnerships"
}

type CustomerPartnershipSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewCustomerPartnershipSeeder(cfg gorm_seeder.SeederConfiguration) CustomerPartnershipSeeder {
	return CustomerPartnershipSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

func (c *CustomerPartnershipSeeder) Seed(db *gorm.DB) error {
	partnerships := make([]CustomerPartnership, 0)

	b2b := CustomerPartnership{
		ID:          "dbe70bc4-89a1-4b20-9c77-18c011d47892",
		Name:        "Business to Business (B2B)",
		Description: "Companies that sell products or services to other businesses rather than directly to consumers. Examples include wholesale distributors, manufacturers, and software providers.",
	}

	partnerships = append(partnerships, b2b)

	b2c := CustomerPartnership{
		ID:          "1501d3a7-11f4-46b1-a9c5-1657f4ff10a1",
		Name:        "Business to Customer (B2C)",
		Description: "Companies that sell products or services directly to individual consumers. Examples include retailers, restaurants, and entertainment services",
	}

	partnerships = append(partnerships, b2c)

	serviceProvider := CustomerPartnership{
		ID:          "12d41010-09ef-4546-a70a-4ef5c76713e6",
		Name:        "Service Provider",
		Description: "Companies that provide services to other businesses or individuals. Examples include consulting firms, marketing agencies, and software development companies.",
	}

	partnerships = append(partnerships, serviceProvider)

	client := CustomerPartnership{
		ID:          "a6f5a1eb-5f65-459d-bc89-d652af7e83f1",
		Name:        "Client",
		Description: "Companies that purchase products or services from other businesses. Examples include manufacturers, retailers, and service providers.",
	}

	partnerships = append(partnerships, client)

	for _, partnership := range partnerships {
		db.FirstOrCreate(&partnership)
	}

	return nil
}

func (c *CustomerTypeSeeder) Clear(db *gorm.DB) error {
	entity := CustomerType{}
	return c.SeederAbstract.Delete(db, entity.TableName())
}
