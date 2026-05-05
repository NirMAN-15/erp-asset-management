package main

import (
    "os"
    "log"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/NirMAN-15/erp-asset-management/backend/config"
    "github.com/NirMAN-15/erp-asset-management/backend/models"
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

    r := gin.Default()
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })
    port := os.Getenv("PORT")
    if port == "" { port = "8080" }
    r.Run(":" + port)
}