package main

import (
	"backend/models"
	"backend/routes" // Import routes package
	"log"
)

func main() {
	// Connect to the database
	models.ConnectDatabase()

	// Set up the routes
	router := routes.SetupRouter() // Call SetupRouter from routes package

	log.Println("Server is running on http://localhost:8080")
	router.Run(":8080")
}
