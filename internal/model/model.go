package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("Failed to connect to database!")
	}
	database.LogMode(true)
	DB = database
}

func RegisterModel(database *gorm.DB, model ...interface{}) *gorm.DB {
	return database.AutoMigrate(model)
}
