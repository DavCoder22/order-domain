package main

import (
	"context"
	"log"

	"order-domain/order-service/src/config"
	"order-domain/order-service/src/handlers"
	"order-domain/order-service/src/repository"
	service "order-domain/order-service/src/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Conexión a la base de datos
func connectDB(cfg *config.Config) (*pgxpool.Pool, error) {
	connStr := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Configurar rutas
func setupRouter(orderHandler *handlers.OrderHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)
	return r
}

func main() {
	// Cargar configuración
	if err := config.LoadConfig("."); err != nil {
		log.Fatalf("❌ Error cargando configuración: %v", err)
	}
	cfg := config.AppConfig

	// Conectar a la BD
	db, err := connectDB(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inicializar servicios y handlers
	repo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(repo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Configurar servidor
	r := setupRouter(orderHandler)

	// Iniciar servidor
	log.Printf("🚀 Servidor iniciado en http://localhost:%s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("❌ Error iniciando servidor: %v", err)
	}
}
