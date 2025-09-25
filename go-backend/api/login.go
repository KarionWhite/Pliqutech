package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginHandler hat jetzt die korrekte Signatur für Gin
func LoginHandler(c *gin.Context) {
	var payload LoginPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Anfrage"})
		return
	}

	var storedHash, username string
	var id int

	err := DB.QueryRow("SELECT id, username, password_hash FROM users WHERE username = ?", payload.Username).Scan(&id, &username, &storedHash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ungültige Anmeldedaten"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ungültige Anmeldedaten"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Login erfolgreich",
		"username": username,
		"userId":   id,
	})
}
