package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey string

// Establece la clave secreta desde el .env, inicializado en main
func SetSecretKey (key string) {
	secretKey = key
}

// Middleware para validar tokens JWT
func AuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		// Obtiene el encabezado de autorización
		authHeader := strings.TrimSpace(ctx.GetHeader("Authorization"))
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		// Divide el encabezado en el Bearer y el Token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			ctx.Abort()
			return
		}

		// Verifica y decodifica el token JWT
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		// Valida el token
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		// Establece los claims del token en el contexto
		ctx.Set("paramedicoId", claims["id"])
		ctx.Set("isAdmin", claims["isAdmin"])
		ctx.Next()
	}
}

// Middleware para verificar si el usuario es admin
func IsAdminMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		isAdmin, _ := ctx.Get("isAdmin")
		if !isAdmin.(bool) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}