package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TemporaryAuthMiddleware es un middleware temporal que permite todas las solicitudes sin verificación de autenticación
func TemporaryAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verificar si el encabezado de autenticación temporal está presente
		if c.GetHeader("X-Temporary-Auth") == "allow" {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Acceso no autorizado"})
		}
	}
}
