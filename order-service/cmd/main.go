package main

import (
	"log"
	"order-service/internal/client"
	"order-service/internal/domain"
	"order-service/internal/repository"
	"order-service/internal/transport/http"
	"order-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user password=0000 dbname=order_db port=5431 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Order DB:", err)
	}

	db.AutoMigrate(&domain.Order{})

	paymentClient := client.NewPaymentClient("http://localhost:8081")
	repo := repository.NewOrderRepository(db)
	uc := usecase.NewOrderUseCase(repo, paymentClient)
	handler := http.NewOrderHandler(uc)

	r := gin.Default()

	r.POST("/orders", handler.CreateOrder)

	r.GET("/orders/:id", handler.GetOrder)

	r.PATCH("/orders/:id/cancel", handler.CancelOrder)

	r.GET("/orders/stats", handler.GetStats)

	log.Println("Order Service starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
