package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	// Load configuration from Viper
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatalf("❌ Error loading configuration: %v", err)
	}

	// Try to connect to RDS
	if connectToRDS(config) {
		log.Println("✅ Successful connection to RDS")
	} else {
		// If it fails, connect to Supabase
		if connectToSupabase(config) {
			log.Println("✅ Successful connection to Supabase")
		} else {
			log.Fatal("❌ Could not connect to any database")
		}
	}
}

func connectToRDS(config Config) bool {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=require",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName,
	)

	log.Println("🔍 Connecting to RDS with:", dataSourceName)

	var dbErr error
	DB, dbErr = pgxpool.New(context.Background(), dataSourceName)
	if dbErr != nil {
		log.Printf("❌ Error opening RDS database: %v", dbErr)
		return false
	}

	pingErr := DB.Ping(context.Background())
	if pingErr != nil {
		log.Printf("❌ Error connecting to RDS database: %v", pingErr)
		return false
	}
	return true
}

func connectToSupabase(config Config) bool {
	supabaseURL := config.SupabaseURL
	if supabaseURL == "" {
		log.Fatal("❌ Supabase URL not configured")
	}

	log.Println("🔍 Connecting to Supabase with:", supabaseURL)

	var dbErr error
	DB, dbErr = pgxpool.New(context.Background(), supabaseURL)
	if dbErr != nil {
		log.Printf("❌ Error opening Supabase database: %v", dbErr)
		return false
	}

	pingErr := DB.Ping(context.Background())
	if pingErr != nil {
		log.Printf("❌ Error connecting to Supabase database: %v", pingErr)
		return false
	}
	return true
}
