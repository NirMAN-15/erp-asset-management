package middleware

import (
    "net/http"; "os"; "strings"
    "github.com/gin-gonic/gin"
    jwt "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        h := c.GetHeader("Authorization")
        if !strings.HasPrefix(h, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort(); return
        }
        t, err := jwt.Parse(strings.TrimPrefix(h, "Bearer "),
            func(t *jwt.Token) (interface{}, error) {
                return []byte(os.Getenv("JWT_SECRET")), nil
            })
        if err != nil || !t.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort(); return
        }
        c.Next()
    }
}