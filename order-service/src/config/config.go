package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	AppPort    string `mapstructure:"APP_PORT"`
}

var AppConfig Config

// Cambia la firma de la función para que solo retorne la configuración
func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app") // Archivo esperado: app.env
	viper.SetConfigType("env")

<<<<<<< HEAD
	viper.AutomaticEnv() // Cargar variables de entorno del sistema

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("⚠️ No se pudo leer el archivo de configuración: %s", err)
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Printf("❌ No se pudo decodificar la configuración: %v", err)
		return err
	}

	log.Println("✅ Configuración cargada correctamente")
	return nil
=======
	viper.AutomaticEnv() // Permite que Viper lea variables de entorno
	if err = viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
		return
	}
	if err = viper.Unmarshal(&config); err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return
	}
	return config, nil
>>>>>>> 2e8c4e40ccb4194782651a6cae4a21614992d7c7
}
