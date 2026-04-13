package main

import (
	"log"
	"net"
	"os"

	"order-service/internal/client"
	"order-service/internal/domain"
	"order-service/internal/repository"
	"order-service/internal/usecase"

	grpcTransport "order-service/internal/transport/grpc"
	httpTransport "order-service/internal/transport/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	orderProto "github.com/dannieey/assignment2-generated/order"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Order DB:", err)
	}
	db.AutoMigrate(&domain.Order{})

	paymentAddr := os.Getenv("PAYMENT_SERVICE_ADDR")
	paymentClient, err := client.NewPaymentGRPCClient(paymentAddr)
	if err != nil {
		log.Fatalf("Failed to connect to Payment Service at %s: %v", paymentAddr, err)
	}

	repo := repository.NewOrderRepository(db)
	uc := usecase.NewOrderUseCase(repo, paymentClient)

	go func() {
		grpcPort := os.Getenv("ORDER_GRPC_PORT")
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("failed to listen Order gRPC: %v", err)
		}

		s := grpc.NewServer()
		orderProto.RegisterOrderServiceServer(s, grpcTransport.NewOrderGRPCHandler(uc))

		log.Printf("Order gRPC Streaming Service starting on %s", grpcPort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve Order gRPC: %v", err)
		}
	}()

	handler := httpTransport.NewOrderHandler(uc)
	r := gin.Default()

	r.POST("/orders", handler.CreateOrder)
	r.GET("/orders/:id", handler.GetOrder)
	r.PATCH("/orders/:id/cancel", handler.CancelOrder)
	r.GET("/orders/stats", handler.GetStats)

	httpPort := os.Getenv("HTTP_PORT")
	log.Printf("Order HTTP Service starting on %s", httpPort)
	if err := r.Run(httpPort); err != nil {
		log.Fatal("Failed to run HTTP server:", err)
	}
}
