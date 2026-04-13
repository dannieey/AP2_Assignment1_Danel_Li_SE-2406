package client

import (
	"context"
	"time"

	// Убедись, что путь к сгенерированному пакету верный
	payment "github.com/dannieey/assignment2-generated/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentGRPCClient struct {
	client payment.PaymentServiceClient
}

// NewPaymentGRPCClient создает новое gRPC соединение
func NewPaymentGRPCClient(addr string) (*PaymentGRPCClient, error) {
	// В новых версиях gRPC вместо Dial рекомендуется NewClient
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &PaymentGRPCClient{
		client: payment.NewPaymentServiceClient(conn),
	}, nil
}

// CheckPayment реализует твой интерфейс PaymentClient
func (c *PaymentGRPCClient) CheckPayment(orderID string, amount int64) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Формируем gRPC запрос из прото-файла
	req := &payment.PaymentRequest{
		OrderId: orderID,
		Amount:  float64(amount), // Приводим к float64, как в .proto
	}

	// Делаем вызов серверу Payment Service
	resp, err := c.client.ProcessPayment(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}

//package client
//
//import (
//	"bytes"
//	"encoding/json"
//	"errors"
//	"net/http"
//	"time"
//)
//
//type PaymentClient struct {
//	httpClient *http.Client
//	baseURL    string
//}
//
//func NewPaymentClient(baseURL string) *PaymentClient {
//	return &PaymentClient{
//		baseURL: baseURL,
//		httpClient: &http.Client{
//			Timeout: 2 * time.Second,
//		},
//	}
//}
//
//func (c *PaymentClient) CheckPayment(orderID string, amount int64) (string, error) {
//	requestBody, _ := json.Marshal(map[string]interface{}{
//		"order_id": orderID,
//		"amount":   amount,
//	})
//
//	resp, err := c.httpClient.Post(c.baseURL+"/payments", "application/json", bytes.NewBuffer(requestBody))
//	if err != nil {
//		return "Failed", err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
//		return "Failed", errors.New("payment service returned an error")
//	}
//
//	var result struct {
//		Status string `json:"status"`
//	}
//	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
//		return "Failed", err
//	}
//
//	return result.Status, nil
//}
