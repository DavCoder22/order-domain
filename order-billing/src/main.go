package main

import (
	"log"
	"order-billing/src/handlers"
	"order-billing/src/repository"
	"order-billing/src/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Configurar la conexión a la base de datos
	db := repository.NewBillingRepository("your-database-url")

	// Crear el servicio de facturación
	billingService := services.NewBillingService(db)

	// Crear el controlador de facturación
	billingHandler := handlers.NewBillingHandler(billingService)

	// Configurar el router
	router := gin.Default()

	// Definir las rutas
	router.POST("/billing", billingHandler.GenerateInvoice)

	// Iniciar el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
