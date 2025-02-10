package main

import (
	"log"
<<<<<<< HEAD:order-service/src/main.go
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
>>>>>>> 125c66e8f56ca6cc5e6ac090cf8992d7170db73d:order-service/main.go

	"github.com/gin-gonic/gin"
)

func main() {
<<<<<<< HEAD:order-service/src/main.go
	// Cargar configuración
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
=======
	// Cargar la configuración
<<<<<<< HEAD:order-service/main.go
	_, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
=======
	if err := config.LoadConfig("."); err != nil {
		log.Fatal("Error cargando configuración:", err)
>>>>>>> 125c66e8f56ca6cc5e6ac090cf8992d7170db73d:order-service/main.go
>>>>>>> 68ff70cdfd0b5117a652956420e31404fe75439d:order-service/src/main.go
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

<<<<<<< HEAD:order-service/src/main.go
	// Middleware de autenticación básico
	r.Use(func(c *gin.Context) {
		if c.GetHeader("Authorization") != "Bearer "+cfg.JWTSecret {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Acceso no autorizado"})
			return
		}
		c.Next()
	})
=======
	// Usar el middleware temporal de autenticación
	r.Use(middleware.TemporaryAuthMiddleware())
>>>>>>> 125c66e8f56ca6cc5e6ac090cf8992d7170db73d:order-service/main.go

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
