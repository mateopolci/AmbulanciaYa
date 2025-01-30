package middleware

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey string

// Establece la clave secreta desde el .env, inicializado en main
func SetSecretKey (key string) {
	secretKey = key
}

func GetSecretKey() string {
    return secretKey
}

func AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        // Get token from cookie instead of header
        tokenString, err := ctx.Cookie("jwt")
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No authentication cookie found"})
            ctx.Abort()
            return
        }

        // Verify and decode JWT token
        claims := jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(secretKey), nil
        })

        // Validate token
        if err != nil || !token.Valid {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            ctx.Abort()
            return
        }

        // Check token expiration
        if exp, ok := claims["exp"].(float64); ok {
            if time.Now().Unix() > int64(exp) {
                ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
                ctx.Abort()
                return
            }
        }

        // Set claims in context
        ctx.Set("paramedicoId", claims["id"])
        ctx.Set("isAdmin", claims["isAdmin"])
        ctx.Next()
    }
}

// Middleware para verificar si el usuario es admin
func IsAdminMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        isAdmin, exists := ctx.Get("isAdmin") 
        if !exists {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Se requiere autenticación"})
            ctx.Abort()
            return
        }

        if !isAdmin.(bool) {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Se requieren privilegios de administrador"})
            ctx.Abort()
            return
        }
        ctx.Next()
    }
}