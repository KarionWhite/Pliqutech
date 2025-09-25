package api

import (
	"net/http"
	// Importiere fmt für Fehlerformatierung
	"github.com/KarionWhite/pliqutech-go-backend/auth" // Importiere dein auth-Package
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// DB-Variable und RegisterPayload-Struktur werden aus types.go geladen
// var DB *sql.DB
// type RegisterPayload struct { ... }

func RegisterHandel(c *gin.Context) {
	var payload RegisterPayload

	// Bindet die JSON-Anfrage an die Payload-Struktur
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Anfrage", "details": err.Error()})
		return
	}

	// --- 1. Prüfen, ob der Username bereits existiert ---
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", payload.Username).Scan(&count)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Datenbankfehler bei der Benutzerprüfung", "details": err.Error()})
		return
	}
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Benutzername bereits vergeben"})
		return
	}

	// --- 2. Passwort hashen ---
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Hashen des Passworts", "details": err.Error()})
		return
	}

	// --- 3. Neuen Benutzer in die Datenbank einfügen (INSERT) ---
	// Stelle sicher, dass deine 'users'-Tabelle eine AUTO_INCREMENT Spalte für 'id' hat.
	result, err := DB.Exec(
		"INSERT INTO users (username, password_hash, email) VALUES (?, ?, ?)",
		payload.Username,
		string(hashedPassword),
		payload.Email,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Speichern des Benutzers in der Datenbank", "details": err.Error()})
		return
	}

	// --- 4. ID des neu erstellten Benutzers abrufen ---
	// Dies ist notwendig, um den Token zu generieren.
	newUserID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Abrufen der neuen Benutzer-ID", "details": err.Error()})
		return
	}

	// --- 5. Token generieren ---
	// Verwende die ID und den Benutzernamen des neu registrierten Benutzers.
	token, err := auth.GenerateToken(int(newUserID), payload.Username) // Cast zu int
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Generieren des Tokens", "details": err.Error()})
		return
	}

	// --- 6. Erfolgsmeldung mit Token zurückgeben ---
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Registrierung erfolgreich",
		"username": payload.Username,
		"userId":   newUserID, // Die neu generierte Benutzer-ID
		"token":    token,     // Der generierte JWT!
	})
}
