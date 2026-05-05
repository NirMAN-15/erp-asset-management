package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/NirMAN-15/erp-asset-management/backend/handlers"
    "github.com/NirMAN-15/erp-asset-management/backend/middleware"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type")
        if c.Request.Method == "OPTIONS" { c.AbortWithStatus(204); return }
        c.Next()
    })
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })
    api := r.Group("/api/v1")
    api.POST("/auth/login", handlers.Login)
    p := api.Group("/")
    p.Use(middleware.AuthMiddleware())
    {
        p.GET("/assets", handlers.GetAssets)
        p.POST("/assets", handlers.CreateAsset)
        p.GET("/assets/:id", handlers.GetAsset)
        p.PUT("/assets/:id", handlers.UpdateAsset)
        p.DELETE("/assets/:id", handlers.DeleteAsset)
        p.GET("/assets/:id/depreciation", handlers.GetDepreciation)
        p.GET("/assignments", handlers.GetAssignments)
        p.POST("/assignments", handlers.AssignAsset)
        p.PUT("/assignments/:id/return", handlers.ReturnAsset)
        p.GET("/assignments/employee/:eid", handlers.GetByEmployee)
    }
    return r
}
