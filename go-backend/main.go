package main

import (
	"database/sql"
	"fmt"
	"log"
	"os" // Wichtig: os-Paket für Umgebungsvariablen importieren

	"github.com/KarionWhite/pliqutech-go-backend/api"
	"github.com/KarionWhite/pliqutech-go-backend/auth"
	"github.com/KarionWhite/pliqutech-go-backend/middleware"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	Port := os.Getenv("PORT")
	jwt_Secret := os.Getenv("JWT_SECRET")

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if Port == "" {
		Port = ":8787"
	}
	if jwt_Secret == "" {
		log.Fatal("Umgebungsvariable JWT_SECRET ist nicht gesetzt!")
	}
	auth.SecretKey = []byte(jwt_Secret)
	if dbUser == "" || dbPass == "" || dbName == "" {
		log.Fatal("Datenbank-Umgebungsvariablen sind nicht korrekt gesetzt!")
	}
	if dbHost == "" {
		dbHost = "127.0.0.1" // Standardwert, falls nicht gesetzt
	}
	if dbPort == "" {
		dbPort = "3306" // Standardwert, falls nicht gesetzt
	}

	// DSN-String dynamisch zusammenbauen
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// --- Datenbank-Setup ---
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Fehler beim Verbinden zur Datenbank:", err)
	}
	defer db.Close()

	api.DB = db

	r := gin.Default()
	config := cors.DefaultConfig()
	// WICHTIG: Ersetze "http://localhost:5173" durch den genauen Origin deines Vue-Dev-Servers.
	// Wenn du von einer anderen IP aus testest (nicht lokal), musst du diese IP hier angeben.
	// Für die Entwicklung ist es manchmal einfacher, AllowAllOrigins auf true zu setzen,
	// aber das ist NICHT für die Produktion gedacht!
	config.AllowOrigins = []string{"http://localhost:5173"}                             // Dein Vue-Frontend Origin
	config.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}           // Erlaubte HTTP-Methoden
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept"} // Wichtige Header für Anfragen
	config.ExposeHeaders = []string{"Content-Length"}                                   // Header, die dem Browser zugänglich gemacht werden sollen
	config.AllowCredentials = true                                                      // Erlaubt das Senden von Cookies oder Auth-Headern
	r.Use(cors.New(config))                                                             // CORS-Middleware auf den Router anwenden
	r.SetTrustedProxies([]string{"127.0.0.1"})
	apiGroup := r.Group("/api")
	{
		// Public Endpoints (kein Token nötig)
		apiGroup.POST("/login", api.LoginHandler)
		apiGroup.POST("/register", api.RegisterHandel) // Dein Registrierungs-Handler

		// Geschützte Endpunkte (Token nötig)
		// Wende die AuthMiddleware auf diese Gruppe an
		protectedGroup := apiGroup.Group("/")
		protectedGroup.Use(middleware.AuthMiddleware()) // HIER WIRD DIE MIDDLEWARE ANGEWENDET
		{
			//Protected apis
		}
	}
	fmt.Printf("Go-Server mit Gin startet auf http://localhost%s\n", Port)
	r.Run(Port)
}
