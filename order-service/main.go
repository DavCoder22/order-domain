package main

import (
	"context"
	"log"
<<<<<<< HEAD
	"net/http"

	"order-service/src/config"
	"order-service/src/handlers"
	"order-service/src/repository"
	"order-service/src/service"
=======

	"order-domain/order-service/src/config"
	"order-domain/order-service/src/handlers"
	"order-domain/order-service/src/middleware"
	"order-domain/order-service/src/repository"
	service "order-domain/order-service/src/services"
>>>>>>> 74fda1443c8ad46b1221ca1a6008498c9f0213f4

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
<<<<<<< HEAD
	// Cargar configuración
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
	}

	// Conexión a PostgreSQL
	connStr := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName
=======
	// Cargar la configuración
	if err := config.LoadConfig("."); err != nil {
		log.Fatal("Error cargando configuración:", err)
	}

	// Conexión a PostgreSQL
	connStr := "postgres://" + config.AppConfig.DBUser + ":" + config.AppConfig.DBPassword + "@" + config.AppConfig.DBHost + ":" + config.AppConfig.DBPort + "/" + config.AppConfig.DBName
>>>>>>> 74fda1443c8ad46b1221ca1a6008498c9f0213f4
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
<<<<<<< HEAD

=======
<<<<<<< HEAD

	// Usar el middleware temporal de autenticación
	r.Use(middleware.TemporaryAuthMiddleware())
=======
>>>>>>> 74fda1443c8ad46b1221ca1a6008498c9f0213f4
	// Middleware de autenticación básico
	r.Use(func(c *gin.Context) {
		if c.GetHeader("Authorization") != "Bearer "+cfg.JWTSecret {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Acceso no autorizado"})
			return
		}
		c.Next()
	})
<<<<<<< HEAD
=======
>>>>>>> 2e8c4e40ccb4194782651a6cae4a21614992d7c7
>>>>>>> 74fda1443c8ad46b1221ca1a6008498c9f0213f4

	// Rutas
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)

	// Iniciar servidor
<<<<<<< HEAD
	log.Printf("Servidor iniciado en el puerto %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
=======
	log.Printf("Servidor iniciado en el puerto %s", config.AppConfig.AppPort)
	if err := r.Run(":" + config.AppConfig.AppPort); err != nil {
>>>>>>> 74fda1443c8ad46b1221ca1a6008498c9f0213f4
		log.Fatal("Error iniciando servidor:", err)
	}
}
