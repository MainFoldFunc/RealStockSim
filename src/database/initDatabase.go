package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/MainFoldFunc/RealStockSim/src/structs"
)

var DB *gorm.DB // Global DB variable

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while opening database:", err)
	}

	// AutoMigrate to create/update table structure
	DB.AutoMigrate(&structs.Users{})

	log.Println("Database connected successfully")
}
