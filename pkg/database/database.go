package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var database, _ = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})


func CreateTables() error {
	// Auto-migrating all the tables
	err := database.AutoMigrate(&User{}, &Wallet{})
	if err != nil {
		return err
	}

	database.Preload(clause.Associations).Find(&Wallet{})
	database.Preload(clause.Associations).Find(&User{})

	return nil
}