package main

import (
	"log"
	"os"

	"github.com/NirMAN-15/erp-asset-management/backend/config"
	"github.com/NirMAN-15/erp-asset-management/backend/models"
	"github.com/NirMAN-15/erp-asset-management/backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}
	config.ConnectDatabase()

	// Auto-migrate creates tables automatically
	config.DB.AutoMigrate(
		&models.Asset{},
		&models.Assignment{},
		&models.Maintenance{},
	)
	log.Println("Tables migrated!")

	router := routes.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
