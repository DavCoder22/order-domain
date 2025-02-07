package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

<<<<<<< HEAD
=======
	"github.com/joho/godotenv"
>>>>>>> 2e8c4e40ccb4194782651a6cae4a21614992d7c7
	_ "github.com/lib/pq"
)

var DB *sql.DB

<<<<<<< HEAD
func InitDB() {
	// Cargar la configuración desde Viper
	err := LoadConfig(".")
=======
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
	DB, err = sql.Open("postgres", dataSourceName)
>>>>>>> 2e8c4e40ccb4194782651a6cae4a21614992d7c7
	if err != nil {
		log.Fatalf("❌ Error cargando la configuración: %v", err)
	}
<<<<<<< HEAD

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
=======
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: %q", err)
>>>>>>> 2e8c4e40ccb4194782651a6cae4a21614992d7c7
	}

	// Probar conexión
	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatalf("❌ Error conectando a la base de datos: %v", pingErr)
	}

	log.Println("✅ Conexión exitosa a la base de datos!")
}
