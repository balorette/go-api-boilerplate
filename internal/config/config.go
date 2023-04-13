package config

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"go-gin-api/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect to db")
	}

	database.AutoMigrate(&models.Book{})

	DB = database
}