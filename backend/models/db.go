package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	// No need to use fmt.Sprintf when the string is static
	dsn := "host=localhost user=user password=password dbname=slackclone port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Run migrations for your models
	database.AutoMigrate(&User{}, &Workspace{}, &Channel{}, &Message{})

	// Assign the DB connection to the global variable
	DB = database
	log.Println("Database connection established")
}
