package models

import (
	"time"

	"github.com/google/uuid"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

type CustomerPartnership struct {
	ID 			string `gorm:"primaryKey" json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`json:"-"`
}

func (c *CustomerPartnership) BeforeCreate() (error) {
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
	partnerships := make([]CustomerPartnership, 0);

	b2b := CustomerPartnership{
		ID: "6b4e49de-431f-49f0-b8b3-c7dd75f20729",
		Name: "Business to Business (B2B)",
		Description: "Companies that sell products or services to other businesses rather than directly to consumers. Examples include wholesale distributors, manufacturers, and software providers.",
	}

	partnerships = append(partnerships, b2b)

	b2c := CustomerPartnership{
		ID: "a1048e79-4cd8-4ff7-9e2f-e4ce7d6a0a8f",
		Name: "Business to Customer (B2C)",
		Description: "Companies that sell products or services directly to individual consumers. Examples include retailers, restaurants, and entertainment services",
	}

	partnerships = append(partnerships, b2c)

	serviceProvider := CustomerPartnership{
		ID: "6ff9e384-d2b7-4083-b4c2-0c0126c6f865",
		Name: "Service Provider",
		Description: "Companies that provide services to other businesses or individuals. Examples include consulting firms, marketing agencies, and software development companies.",
	}

	partnerships = append(partnerships, serviceProvider)

	client := CustomerPartnership{
		ID: "a6f5a1eb-5f65-459d-bc89-d652af7e83f1",
		Name: "Client",
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


