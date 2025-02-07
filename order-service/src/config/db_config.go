package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

<<<<<<< HEAD
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
=======
	_ "github.com/lib/pq"
>>>>>>> 125c66e8f56ca6cc5e6ac090cf8992d7170db73d
)

var DB *sql.DB

<<<<<<< HEAD
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %q", err)
	}

	var dbUser = os.Getenv("DB_USER")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbName = os.Getenv("DB_NAME")

	var dataSourceName = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
=======
func InitDB() {
	// Cargar la configuración desde Viper
	err := LoadConfig(".")
>>>>>>> 125c66e8f56ca6cc5e6ac090cf8992d7170db73d
	if err != nil {
		log.Fatalf("❌ Error cargando la configuración: %v", err)
	}

	// Construcción de la cadena de conexión correcta para PostgreSQL
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=require",
		AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBName,
	)

	log.Println("🔍 Conectando a la base de datos con:", dataSourceName)

	// Abrir conexión
	var dbErr error
	DB, dbErr = sql.Open("postgres", dataSourceName)
	if dbErr != nil {
		log.Fatalf("❌ Error al abrir la base de datos: %v", dbErr)
	}

	// Probar conexión
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatalf("❌ Error conectando a la base de datos: %v", pingErr)
	}

	log.Println("✅ Conexión exitosa a la base de datos!")
}
