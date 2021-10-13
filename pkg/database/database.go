package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database, _ = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})


func CreateTables() error {
	// Auto-migrating all the tables
	err := database.AutoMigrate(&User{})
	if err != nil {
		return err
	}

	return nil
}