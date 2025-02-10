package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	SupabaseURL string `mapstructure:"SUPABASE_URL"`
	AppPort     string `mapstructure:"APP_PORT"`
}

var AppConfig Config

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv() // Allows Viper to read environment variables

	if err = viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return
	}

	AppConfig = config // Assign the global configuration
	return config, nil
}
