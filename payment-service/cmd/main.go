package main

import (
	"log"
	"net"
	"os"
	"payment-service/internal/domain"
	"payment-service/internal/repository"
	"payment-service/internal/usecase"

	// Правильные пути согласно твоей архитектуре (internal/transport/...)
	grpcTransport "payment-service/internal/transport/grpc"
	httpTransport "payment-service/internal/transport/http"

	payment "github.com/dannieey/assignment2-generated/payment"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 1. Загрузка .env
	godotenv.Load()

	// 2. Чтение конфига
	dsn := os.Getenv("DATABASE_URL")
	grpcPort := os.Getenv("GRPC_PORT")
	httpPort := os.Getenv("HTTP_PORT")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&domain.Payment{})

	repo := repository.NewPaymentRepository(db)
	uc := usecase.NewPaymentUseCase(repo)

	// --- gRPC Server ---
	go func() {
		lis, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("failed to listen gRPC: %v", err)
		}

		s := grpc.NewServer()
		// Используем алиас grpcTransport
		payment.RegisterPaymentServiceServer(s, grpcTransport.NewPaymentGRPCHandler(uc))

		log.Printf("gRPC Payment Service starting on :%s", grpcPort)
		s.Serve(lis)
	}()

	// --- HTTP Server (Gin) ---
	// Используем алиас httpTransport
	handler := httpTransport.NewPaymentHandler(uc)
	r := gin.Default()

	r.POST("/payments", handler.CreatePayment)
	r.GET("/payments/:order_id", handler.GetPayment)

	log.Printf("HTTP Payment Service starting on :%s", httpPort)
	r.Run(":" + httpPort)
}
