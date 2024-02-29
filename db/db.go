package db

import (
	"um/models"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=root dbname=users port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Msgf("Failed to connect to database: %v", err)
		err = db.Exec("CREATE DATABASE users").Error
		if err != nil {
			panic("failed to create database")
		}
	}
	log.Info().Msg("Connected to db...")
	// Create the table
	db.AutoMigrate(&models.User{})
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
