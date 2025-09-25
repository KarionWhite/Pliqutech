package middleware

import (
	"net/http"

	"github.com/KarionWhite/pliqutech-go-backend/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header fehlt"})
			return
		}

		tokenString, err := auth.ExtractTokenFromHeader(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Ungültiger oder abgelaufener Token", "details": err.Error()})
			return
		}

		// Den UserID und Username in den Gin-Kontext speichern,
		// damit nachfolgende Handler darauf zugreifen können.
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next() // Anfrage an den nächsten Handler in der Kette weiterleiten
	}
}
