package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Cargar la configuración desde Viper
	err := LoadConfig(".")
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
