package main

import (
	"log"

	"order-domain/order-service/src/config"
	"order-domain/order-service/src/handlers"
	"order-domain/order-service/src/middleware"
	"order-domain/order-service/src/repository"
	service "order-domain/order-service/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar la configuración
	_, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
	}

	// Inicializar la conexión a la base de datos
	config.InitDB()

	// Usar la conexión global inicializada en config.InitDB()
	db := config.DB
	defer db.Close()

	// Inicializar servicios
	repo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(repo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Configurar servidor Gin
	r := gin.Default()

	// Usar el middleware temporal de autenticación
	r.Use(middleware.TemporaryAuthMiddleware())

	// Definir rutas
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)

	// Iniciar servidor
	log.Printf("Servidor iniciado en el puerto %s", config.AppConfig.AppPort)
	if err := r.Run(":" + config.AppConfig.AppPort); err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}
