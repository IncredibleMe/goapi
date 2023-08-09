package models

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Movie{})
	if err != nil {
		return
	}
	err = database.AutoMigrate(&Thema{})
	if err != nil {
		return
	}

	DB = database
}
