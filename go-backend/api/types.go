package api

import "database/sql"

// RegisterPayload ist die Struktur für die Registrierungsanfrage vom Frontend.
type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// LoginPayload ist die Struktur für die Login-Anfrage vom Frontend.
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DB ist die globale Datenbankverbindung für das api-Package.
// Sie muss in main.go initialisiert und hier zugewiesen werden.
var DB *sql.DB