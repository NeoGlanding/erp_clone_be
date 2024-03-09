package models

import (
	"time"

	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

type CustomerType struct {
	ID 			string `gorm:"primaryKey;unique" json:"id"`

	Name        string `json:"name"`
	Description string `json:"description"`

	Customer []Customer `gorm:"foreignKey:CustomerTypeId" json:"customer"`

	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`json:"-"`
}

func (c *CustomerType) TableName() string {
	return "customer_types"
}

type CustomerTypeSeeder struct {
	gorm_seeder.SeederAbstract
}

func NewCustomerTypeSeeder(cfg gorm_seeder.SeederConfiguration) CustomerTypeSeeder {
	return CustomerTypeSeeder{gorm_seeder.NewSeederAbstract(cfg)}
}

func (s *CustomerTypeSeeder) Seed(db *gorm.DB) error {
	// Create a slice to store instances
	customerTypes := make([]CustomerType, 0)

	// Define instances with static UUIDs and append them to the slice
	soleProprietorship := CustomerType{
		ID:          "6b4e49de-431f-49f0-b8b3-c7dd75f20729",
		Name:        "Sole Proprietorship",
		Description: "A business owned and operated by a single individual. The owner is personally responsible for all business aspects, including debts and liabilities.",
	}
	customerTypes = append(customerTypes, soleProprietorship)

	llc := CustomerType{
		ID:          "a1048e79-4cd8-4ff7-9e2f-e4ce7d6a0a8f",
		Name:        "Limited Liability Company (LLC)",
		Description: "A flexible form of business organization that combines elements of both a corporation and a partnership. Owners have limited liability, and there is flexibility in management and taxation.",
	}
	customerTypes = append(customerTypes, llc)

	privateLtd := CustomerType{
		ID:          "6ff9e384-d2b7-4083-b4c2-0c0126c6f865",
		Name:        "Private Limited Company",
		Description: "A privately held business entity where shares of the company are held by a few individuals, and there are restrictions on share transfer. Owners have limited liability.",
	}
	customerTypes = append(customerTypes, privateLtd)

	publicLtd := CustomerType{
		ID:          "a6f5a1eb-5f65-459d-bc89-d652af7e83f1",
		Name:        "Public Limited Company",
		Description: "A company whose shares are traded freely on a stock exchange. It can raise capital from the public by issuing shares.",
	}
	customerTypes = append(customerTypes, publicLtd)

	holdingCompany := CustomerType{
		ID:          "d7879e88-1f9a-495d-aa6d-4d205e8a5e84",
		Name:        "Holding Company",
		Description: "A company that doesn't engage in active business operations itself but owns a significant amount of voting stock in other companies (subsidiaries). Its purpose is often to control and manage other companies.",
	}
	customerTypes = append(customerTypes, holdingCompany)

	subSidiaryCompany := CustomerType{
		ID:          "c8f22c7b-8a77-48b2-bfb2-155af90b92b8",
		Name:        "Subsidiary Company",
		Description: "A company that is owned or controlled by another company, known as the parent company or holding company.",
	}
	customerTypes = append(customerTypes, subSidiaryCompany)

	nonprofit := CustomerType{
		ID:          "aaf7b3a3-8cc9-4f92-85b8-cd9b7eab39ad",
		Name:        "Nonprofit Organization",
		Description: "An organization formed for purposes other than making a profit. Nonprofits often pursue charitable, educational, or social goals.",
	}
	customerTypes = append(customerTypes, nonprofit)

	cooperative := CustomerType{
		ID:          "9c332de2-50b8-4341-858b-e7925fb1392a",
		Name:        "Cooperative",
		Description: "An organization owned and operated by a group of individuals for their mutual benefit. Members share profits or benefits based on their contributions.",
	}
	customerTypes = append(customerTypes, cooperative)

	franchise := CustomerType{
		ID:          "c23e1d57-96fc-4314-ba92-35a6e5aaf280",
		Name:        "Franchise",
		Description: "A business model where an individual or entity (franchisee) pays fees to the owner (franchisor) for the right to operate under the franchisor's brand and business model.",
	}
	customerTypes = append(customerTypes, franchise)

	jointVenture := CustomerType{
		ID:          "fd50d861-95ee-491d-837b-d95e71987f46",
		Name:        "Joint Venture",
		Description: "A business arrangement where two or more parties agree to pool their resources for a specific project or business activity. The joint venture is a separate entity.",
	}
	customerTypes = append(customerTypes, jointVenture)

	llp := CustomerType{
		ID:          "d163495d-8832-43c0-b6ee-c3a71f71c0e1",
		Name:        "Limited Liability Partnership (LLP)",
		Description: "A partnership in which some or all partners have limited liability. It provides flexibility in management while offering limited liability protection.",
	}
	customerTypes = append(customerTypes, llp)

	cooperativeSociety := CustomerType{
		ID:          "2c0c5f11-5f5c-496c-8f50-70882d367cb4",
		Name:        "Cooperative Society",
		Description: "A form of organization where individuals voluntarily join together to meet common economic, social, and cultural needs through a jointly-owned and democratically-controlled enterprise.",
	}
	customerTypes = append(customerTypes, cooperativeSociety)

	startup := CustomerType{
		ID:          "b25fb049-9d07-4a76-9391-7c2aee0f6664",
		Name:        "Startup",
		Description: "A newly established business with high growth potential, typically in the technology or innovation sector.",
	}
	customerTypes = append(customerTypes, startup)

	for _, element := range customerTypes {
		db.FirstOrCreate(&element)
	}
	
	return nil
	
}

func (c *CustomerPartnershipSeeder) Clear(db *gorm.DB) error {
    entity := CustomerType{}
    return c.SeederAbstract.Delete(db, entity.TableName())
}