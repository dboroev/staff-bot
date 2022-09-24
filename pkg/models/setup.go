package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(postgres.Open("postgres://staff-bot:staff-bot@postgresdb:5432/staff-bot"))
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Job{})
	db.AutoMigrate(&Point{})
	db.AutoMigrate(&Claim{})
	db.AutoMigrate(&UserClaim{})

	DB = db
}
