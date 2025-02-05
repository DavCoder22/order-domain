package main

import (
	"context"
	"log"

	"order-domain/order-service/src/config"
	"order-domain/order-service/src/handlers"
	"order-domain/order-service/src/middleware"
	"order-domain/order-service/src/repository"
	service "order-domain/order-service/src/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Cargar la configuración
	if err := config.LoadConfig("."); err != nil {
		log.Fatal("Error cargando configuración:", err)
	}

	// Conexión a PostgreSQL
	connStr := "postgres://" + config.AppConfig.DBUser + ":" + config.AppConfig.DBPassword + "@" + config.AppConfig.DBHost + ":" + config.AppConfig.DBPort + "/" + config.AppConfig.DBName
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Error conectando a PostgreSQL:", err)
	}
	defer db.Close()

	// Inicializar servicios
	repo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(repo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Configurar servidor Gin
	r := gin.Default()

	// Usar el middleware temporal de autenticación
	r.Use(middleware.TemporaryAuthMiddleware())

	// Rutas
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)

	// Iniciar servidor
	log.Printf("Servidor iniciado en el puerto %s", config.AppConfig.AppPort)
	if err := r.Run(":" + config.AppConfig.AppPort); err != nil {
		log.Fatal("Error iniciando servidor:", err)
	}
}
