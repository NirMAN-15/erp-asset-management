package handlers

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/NirMAN-15/erp-asset-management/backend/config"
    "github.com/NirMAN-15/erp-asset-management/backend/models"
)

func GetAssignments(c *gin.Context) {
    var list []models.Assignment
    config.DB.Where("returned_at IS NULL").Find(&list)
    c.JSON(http.StatusOK, gin.H{"data": list})
}

func AssignAsset(c *gin.Context) {
    var a models.Assignment
    c.ShouldBindJSON(&a)
    a.ID = uuid.NewString()
    a.AssignedAt = time.Now()
    config.DB.Create(&a)
    config.DB.Model(&models.Asset{}).
        Where("id = ?", a.AssetID).Update("status", "ASSIGNED")
    c.JSON(http.StatusCreated, gin.H{"data": a})
}

func ReturnAsset(c *gin.Context) {
    var a models.Assignment
    config.DB.First(&a, "id = ?", c.Param("id"))
    now := time.Now()
    a.ReturnedAt = &now
    config.DB.Save(&a)
    config.DB.Model(&models.Asset{}).
        Where("id = ?", a.AssetID).Update("status", "ACTIVE")
    c.JSON(http.StatusOK, gin.H{"message": "Asset returned"})
}

func GetByEmployee(c *gin.Context) {
    var list []models.Assignment
    config.DB.Where("employee_id = ? AND returned_at IS NULL", c.Param("eid")).Find(&list)
    c.JSON(http.StatusOK, gin.H{"data": list})
}
