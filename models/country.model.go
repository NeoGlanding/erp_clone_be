package models

import (
	"fmt"
	"time"

	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

type Country struct {
	ID				string		`json:"id" gorm:"primaryKey"`
	Name			string		`json:"name"`
	TwoAlphaCode	string		`json:"two_alpha_code"`
	ThreeAlphaCode	string		`json:"three_alpha_code"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		*time.Time	`json:"updated_at"`
	DeletedAt		*time.Time	
}

func (c *Country) TableName() string {
	return "countries"
}

type CountrySeeder struct {
	gorm_seeder.SeederAbstract
}

func NewCountrySeeder(cofg gorm_seeder.SeederConfiguration) CountrySeeder {
	return CountrySeeder{gorm_seeder.NewSeederAbstract(cofg)}
}

func (s *CountrySeeder) Seed(db *gorm.DB) error {
	fmt.Println("Seeding Countries")
	
	indonesia := Country{
		ID: "a65e1a7e-3261-4c7f-95a8-aaaba7ad1998",
		Name: "Indonesia",
		TwoAlphaCode: "ID",
		ThreeAlphaCode: "IDN",
	};

	malaysia := Country{
		ID: "fe18a5d7-78ce-451f-a801-6c4fe53801a2",
		Name: "Malaysia",
		TwoAlphaCode: "MY",
		ThreeAlphaCode: "MYS",
	}

	singapore := Country{
		ID: "879548d7-62ae-4aac-8ab0-6f9f23d7bd64",
		Name: "Singapore",
		TwoAlphaCode: "SG",
		ThreeAlphaCode: "SGP",
	}

	var countries []Country = make([]Country, 3)

	countries[0] = singapore
	countries[1] = malaysia
	countries[2] = indonesia

	return db.CreateInBatches(countries, 3).Error
}

func (c *CountrySeeder) Clear(db *gorm.DB) error {
	entity := Country{}
	return c.SeederAbstract.Delete(db, entity.TableName())
}