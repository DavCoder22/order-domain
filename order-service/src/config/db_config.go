package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	// Cargar la configuración desde Viper
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatalf("❌ Error cargando la configuración: %v", err)
	}

	// Intentar conectar a RDS
	if connectToRDS(config) {
		log.Println("✅ Conexión exitosa a RDS")
	} else {
		// Si falla, conectar a Supabase
		if connectToSupabase(config) {
			log.Println("✅ Conexión exitosa a Supabase")
		} else {
			log.Fatal("❌ No se pudo conectar a ninguna base de datos")
		}
	}
}

func connectToRDS(config Config) bool {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=require",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName,
	)

	log.Println("🔍 Conectando a RDS con:", dataSourceName)

	var dbErr error
	DB, dbErr = pgxpool.New(context.Background(), dataSourceName)
	if dbErr != nil {
		log.Printf("❌ Error al abrir la base de datos RDS: %v", dbErr)
		return false
	}

	pingErr := DB.Ping(context.Background())
	if pingErr != nil {
		log.Printf("❌ Error conectando a la base de datos RDS: %v", pingErr)
		return false
	}
	return true
}

func connectToSupabase(config Config) bool {
	supabaseURL := config.SupabaseURL
	if supabaseURL == "" {
		log.Fatal("❌ URL de Supabase no configurada")
	}

	log.Println("🔍 Conectando a Supabase con:", supabaseURL)

	var dbErr error
	DB, dbErr = pgxpool.New(context.Background(), supabaseURL)
	if dbErr != nil {
		log.Printf("❌ Error al abrir la base de datos Supabase: %v", dbErr)
		return false
	}

	pingErr := DB.Ping(context.Background())
	if pingErr != nil {
		log.Printf("❌ Error conectando a la base de datos Supabase: %v", pingErr)
		return false
	}
	return true
}
