package handlers

import (
	"github.com/NirMAN-15/erp-asset-management/backend/config"
	"github.com/NirMAN-15/erp-asset-management/backend/models"
	"github.com/gin-gonic/gin"
)

func GetAssets(c *gin.Context) {
	var assets []models.Asset
	q := config.DB
	if s := c.Query("status"); s != "" {
		q = q.Where("status = ?", s)
	}
	if cat := c.Query("category"); cat != "" {
		q = q.Where("category = ?", cat)
	}
	q.Order("created_at desc").Find(&assets)
	c.JSON(200, gin.H{"data": assets, "count": len(assets)})
}

func CreateAsset(c *gin.Context) {
	var a models.Asset
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	a.CurrentValue = a.PurchaseValue
	if err := config.DB.Create(&a).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"data": a})
}

func GetAsset(c *gin.Context) {
	var a models.Asset
	if err := config.DB.First(&a, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	c.JSON(200, gin.H{"data": a})
}

func UpdateAsset(c *gin.Context) {
	var a models.Asset
	if err := config.DB.First(&a, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	c.ShouldBindJSON(&a)
	config.DB.Save(&a)
	c.JSON(200, gin.H{"data": a})
}

func DeleteAsset(c *gin.Context) {
	config.DB.Model(&models.Asset{}).
		Where("id = ?", c.Param("id")).
		Update("status", "DISPOSED")
	c.JSON(200, gin.H{"message": "Asset disposed"})
}

func GetDepreciation(c *gin.Context) {
	var a models.Asset
	if err := config.DB.First(&a, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	// Straight-line: 5 year useful life
	annualDep := a.PurchaseValue / 5
	c.JSON(200, gin.H{
		"asset_id":            a.ID,
		"purchase_value":      a.PurchaseValue,
		"annual_depreciation": annualDep,
		"current_value":       a.CurrentValue,
		"method":              "STRAIGHT_LINE",
	})
}
