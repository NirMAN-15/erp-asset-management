package handlers

import (
"net/http"
"time"

"github.com/gin-gonic/gin"
"github.com/google/uuid"
"github.com/NirMAN-15/erp-asset-management/backend/config"
"github.com/NirMAN-15/erp-asset-management/backend/models"
)

func GetMaintenance(c *gin.Context) {
var list []models.Maintenance
config.DB.Order("scheduled_date desc").Find(&list)
c.JSON(http.StatusOK, gin.H{"data": list})
}

func CreateMaintenance(c *gin.Context) {
var m models.Maintenance
if err := c.ShouldBindJSON(&m); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}
m.ID = uuid.NewString()
m.Status = "SCHEDULED"
if err := config.DB.Create(&m).Error; err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}
c.JSON(http.StatusCreated, gin.H{"data": m})
}

func CompleteMaintenance(c *gin.Context) {
var m models.Maintenance
if err := config.DB.First(&m, "id = ?", c.Param("id")).Error; err != nil {
c.JSON(http.StatusNotFound, gin.H{"error": "Maintenance record not found"})
return
}
now := time.Now()
m.CompletedDate = &now
m.Status = "COMPLETED"
config.DB.Save(&m)
c.JSON(http.StatusOK, gin.H{"message": "Maintenance completed", "data": m})
}
