package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// SecretKey ist der geheime Schlüssel, der zum Signieren und Überprüfen der JWTs verwendet wird.
// WICHTIG: Dies sollte ein sehr starker, zufälliger String sein und NICHT im Code hartkodiert werden!
// Am besten über eine Umgebungsvariable laden.
var SecretKey = []byte("super_geheimer_schluessel_bitte_aendern_und_sicher_halten") // Beispiel

// Claims sind die Daten, die in deinem JWT gespeichert werden sollen.
// Wir erweitern jwt.RegisteredClaims um benutzerdefinierte Felder.
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken generiert einen neuen JWT für den gegebenen Benutzer.
func GenerateToken(userID int, username string) (string, error) {
	// Token läuft in 24 Stunden ab
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Ablaufzeitpunkt
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // Ausstellungszeitpunkt
			NotBefore: jwt.NewNumericDate(time.Now()),     // Gültig ab jetzt
			Issuer:    "pliqutech-backend",                // Wer den Token ausgestellt hat
			Subject:   fmt.Sprintf("%d", userID),          // Betreff (oft die Benutzer-ID)
			Audience:  jwt.ClaimStrings{"web"},            // Zielgruppe
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // HS256 ist ein gängiger Signaturalgorithmus
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", fmt.Errorf("Fehler beim Signieren des Tokens: %w", err)
	}
	return tokenString, nil
}

// ValidateToken validiert einen JWT und gibt die Claims zurück, falls gültig.
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Überprüfe den Signaturalgorithmus
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unerwarteter Signaturalgorithmus: %v", token.Header["alg"])
		}
		return SecretKey, nil // Geheimen Schlüssel zum Überprüfen der Signatur zurückgeben
	})

	if err != nil {
		return nil, fmt.Errorf("Fehler beim Parsen/Validieren des Tokens: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token ist ungültig")
	}

	// Prüfe, ob die Claims vom richtigen Typ sind
	if c, ok := token.Claims.(*Claims); ok {
		return c, nil
	} else {
		return nil, fmt.Errorf("Token-Claims sind nicht vom erwarteten Typ")
	}
}

// ExtractTokenFromHeader extrahiert den JWT aus dem "Authorization: Bearer <token>" Header.
func ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("Authorization-Header fehlt")
	}
	// Erwarte das Format "Bearer <TOKEN>"
	const bearerPrefix = "Bearer "
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		return authHeader[len(bearerPrefix):], nil
	}
	return "", fmt.Errorf("Ungültiges Authorization-Header-Format")
}
