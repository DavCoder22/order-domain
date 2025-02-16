package config

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	SupabaseURL     string
	AppPort         string
	JWTSecret       string
	OrderHistoryDB  string // Ejemplo de configuración específica
	OrderTrackingDB string // Ejemplo de configuración específica
}

var AppConfig Config
var DB *pgxpool.Pool
var MySQLDB *sql.DB

func LoadConfig(path string) error {
	// Load environment variables from app.env file
	err := godotenv.Load(path + "/app.env")
	if err != nil {
		log.Printf("Error loading app.env file: %v", err)
	}

	viper.AutomaticEnv() // Allows Viper to read environment variables

	// Set up Viper to read from environment variables
	viper.BindEnv("SUPABASE_URL", "SUPABASE_URL")
	viper.BindEnv("APP_PORT", "APP_PORT")
	viper.BindEnv("JWT_SECRET", "JWT_SECRET")
	viper.BindEnv("MYSQL_DSN", "MYSQL_DSN")

	AppConfig.SupabaseURL = viper.GetString("SUPABASE_URL")
	AppConfig.AppPort = viper.GetString("APP_PORT")
	AppConfig.JWTSecret = viper.GetString("JWT_SECRET")

	return nil
}

func InitDB() {
	supabaseURL := os.Getenv("SUPABASE_URL")
	if supabaseURL == "" {
		log.Fatal("❌ Supabase URL not configured")
	}

	log.Println("🔍 Connecting to Supabase with:", supabaseURL)

	var dbErr error
	DB, dbErr = pgxpool.New(context.Background(), supabaseURL)
	if dbErr != nil {
		log.Fatalf("❌ Error opening Supabase database: %v", dbErr)
	}

	pingErr := DB.Ping(context.Background())
	if pingErr != nil {
		log.Fatalf("❌ Error connecting to Supabase database: %v", pingErr)
	}

	log.Println("✅ Successful connection to Supabase")

	// Conectar a MySQL
	mysqlDSN := os.Getenv("MYSQL_DSN")
	if mysqlDSN == "" {
		log.Fatal("❌ MySQL DSN not configured")
	}

	var mysqlErr error
	MySQLDB, mysqlErr = sql.Open("mysql", mysqlDSN)
	if mysqlErr != nil {
		log.Fatalf("❌ Error opening MySQL database: %v", mysqlErr)
	}

	mysqlPingErr := MySQLDB.Ping()
	if mysqlPingErr != nil {
		log.Fatalf("❌ Error connecting to MySQL database: %v", mysqlPingErr)
	}

	log.Println("✅ Successful connection to MySQL")
}
