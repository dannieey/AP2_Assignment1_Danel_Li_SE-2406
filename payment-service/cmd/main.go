package main

import (
	"log"
	"payment-service/internal/domain"
	"payment-service/internal/repository"
	"payment-service/internal/transport/http"
	"payment-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=user password=0000 dbname=payment_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&domain.Payment{})

	repo := repository.NewPaymentRepository(db)
	uc := usecase.NewPaymentUseCase(repo)
	handler := http.NewPaymentHandler(uc)

	r := gin.Default()

	r.POST("/payments", handler.CreatePayment)
	r.GET("/payments/:order_id", handler.GetPayment)

	log.Println("Payment Service starting on :8081")
	r.Run(":8081")
}
