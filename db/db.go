package db

import (
	// "log"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("postgres", "host=myhost port=myport user=myuser dbname=mydbname password=mypassword sslmode=disable")
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
	}
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
