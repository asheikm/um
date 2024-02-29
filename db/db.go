package db

import (
	"fmt"
	"um/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=root dbname=users port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
		panic("Failed to connect to database")
	}

	// Create the table
	db.AutoMigrate(&models.User{})
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
