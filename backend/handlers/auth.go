package handlers

import (
    "net/http"; "os"; "time"
    "github.com/gin-gonic/gin"
    jwt "github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Login(c *gin.Context) {
    var input LoginInput
    c.ShouldBindJSON(&input)

    // Simple check — in production use DB + bcrypt
    if input.Username != "admin" || input.Password != "admin123" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user": input.Username,
        "exp":  time.Now().Add(24 * time.Hour).Unix(),
    })
    tokenStr, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    c.JSON(http.StatusOK, gin.H{"token": tokenStr, "user": input.Username})
}