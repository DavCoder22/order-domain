package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables")
	}

	// Obtener variables de entorno
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Formatear la cadena de conexión correctamente para PostgreSQL
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Abrir la conexión con la base de datos
	var dbErr error
	DB, dbErr = sql.Open("postgres", dataSourceName)
	if dbErr != nil {
		log.Fatalf("❌ Error opening database: %v", dbErr)
	}

	// Probar la conexión
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatalf("❌ Error connecting to the database: %v", pingErr)
	}

	log.Println("✅ Successfully connected to the database!")
}
