package grpc

import (
	"context"
	"time"

	"github.com/dannieey/assignment2-generated/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentGRPCClient struct {
	client payment.PaymentServiceClient
}

func NewPaymentGRPCClient(addr string) (*PaymentGRPCClient, error) {
	// Устанавливаем соединение с Payment Service
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &PaymentGRPCClient{
		client: payment.NewPaymentServiceClient(conn),
	}, nil
}

func (c *PaymentGRPCClient) CheckPayment(orderID string, amount int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Формируем gRPC запрос
	req := &payment.PaymentRequest{
		OrderId: orderID,
		Amount:  float64(amount), // Конвертируем для соответствия proto-файлу
	}

	// Вызываем удаленный метод ProcessPayment в другом сервисе
	resp, err := c.client.ProcessPayment(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}
