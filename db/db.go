package db

import (
	// "log"
	"fmt"

	"um/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // For now we are not going to use postgres
	// _ "github.com/mattn/go-sqlite3"
	//""
	// "um/db"
	// "gorm.io/driver/sqlite"
	//"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost	 port=5432 user=postgres dbname=users password=root sslmode=disable")
	// do this for sqlite
	// db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
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
