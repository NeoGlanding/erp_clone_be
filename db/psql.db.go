package db

import (
	"fmt"

	"github.com/automa8e_clone/models"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

var PSQL *gorm.DB

func PSQLMigrate() {
	PSQL.AutoMigrate(&models.User{})
	PSQL.AutoMigrate(&models.UserDetails{})
	PSQL.AutoMigrate(&models.Country{})
	PSQL.AutoMigrate(&models.Party{})
	PSQL.AutoMigrate(&models.UserPartyPermission{})
	PSQL.AutoMigrate(&models.File{})
	PSQL.AutoMigrate(&models.CustomerType{})
	PSQL.AutoMigrate(&models.CustomerPartnership{})
	PSQL.AutoMigrate(&models.Customer{})
	PSQL.AutoMigrate(&models.CustomerAddresses{})
}

func PSQLSeed() {
	// Country
	countrySeeder := models.NewCountrySeeder(gorm_seeder.SeederConfiguration{Rows: 3})
	seedersStack := gorm_seeder.NewSeedersStack(PSQL)
	seedersStack.AddSeeder(&countrySeeder)
	err := seedersStack.Seed()
	if err != nil {
		fmt.Println(err)
	}

	// Customer Types
	customerTypesSeeder := models.NewCustomerTypeSeeder(gorm_seeder.SeederConfiguration{Rows: 14})
	ctStack := gorm_seeder.NewSeedersStack(PSQL)
	ctStack.AddSeeder(&customerTypesSeeder)

	ctStackErr := ctStack.Seed()

	if ctStackErr != nil {
		fmt.Println("No error")
	}

	// Customer Partnership
	customerPartnershipSeeder := models.NewCustomerPartnershipSeeder(gorm_seeder.SeederConfiguration{Rows: 3})
	cpStack := gorm_seeder.NewSeedersStack(PSQL)
	cpStack.AddSeeder(&customerPartnershipSeeder)

	cpStackErr := cpStack.Seed()

	if cpStackErr != nil {
		fmt.Println("No error")
	}

}
