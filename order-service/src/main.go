package main

import (
	"context"
	"log"
	"net/http"

	"order-service/src/config"
	"order-service/src/handlers"
	"order-service/src/repository"
	"order-service/src/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Cargar configuración
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
	}

	// Conexión a PostgreSQL
	connStr := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName
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

	// Middleware de autenticación básico
	r.Use(func(c *gin.Context) {
		if c.GetHeader("Authorization") != "Bearer "+cfg.JWTSecret {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Acceso no autorizado"})
			return
		}
		c.Next()
	})

	// Rutas
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id/status", orderHandler.UpdateOrderStatus)

	// Iniciar servidor
	log.Printf("Servidor iniciado en el puerto %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Error iniciando servidor:", err)
	}
}
