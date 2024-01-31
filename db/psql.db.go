package db

import (
	"fmt"

	"github.com/automa8e_clone/models"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

var PSQL *gorm.DB;

func PSQLMigrate() {
	PSQL.AutoMigrate(&models.User{})
	PSQL.AutoMigrate(&models.UserDetails{})
	PSQL.AutoMigrate(&models.Country{})
	PSQL.AutoMigrate(&models.Party{})
	PSQL.AutoMigrate(&models.UserPartyPermission{})
	PSQL.AutoMigrate(&models.File{})
}

func PSQLSeed() {
	// Country
	countrySeeder := models.NewCountrySeeder(gorm_seeder.SeederConfiguration{Rows: 3})
	seedersStack := gorm_seeder.NewSeedersStack(PSQL)
	seedersStack.AddSeeder(&countrySeeder)
	err := seedersStack.Seed()
	if (err != nil) {
		fmt.Println(err)
	}


}