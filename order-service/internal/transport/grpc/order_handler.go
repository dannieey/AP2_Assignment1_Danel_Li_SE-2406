package grpc

import (
	"log"
	"order-service/internal/usecase"

	"github.com/dannieey/assignment2-generated/order"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderGRPCHandler struct {
	order.UnimplementedOrderServiceServer
	usecase *usecase.OrderUseCase
}

func NewOrderGRPCHandler(uc *usecase.OrderUseCase) *OrderGRPCHandler {
	return &OrderGRPCHandler{usecase: uc}
}

func (h *OrderGRPCHandler) SubscribeToOrderUpdates(req *order.OrderSubscriptionRequest, stream grpc.ServerStreamingServer[order.OrderStatusUpdate]) error {
	filterID := req.GetOrderId()
	log.Printf("New subscription started. Filter OrderID: '%s'", filterID)

	updates := h.usecase.GetUpdatesChannel()

	for {
		select {
		case <-stream.Context().Done():
			log.Println("Client disconnected from gRPC stream")
			return nil
		case updatedOrder, ok := <-updates:
			if !ok {
				log.Println("Updates channel closed")
				return nil
			}

			// ГЛАВНОЕ ИСПРАВЛЕНИЕ:
			// Если filterID пустой ("") ИЛИ совпадает с ID заказа — отправляем данные
			if filterID == "" || updatedOrder.ID == filterID {
				log.Printf("Streaming update to client for Order: %s (Status: %s)", updatedOrder.ID, updatedOrder.Status)

				err := stream.Send(&order.OrderStatusUpdate{
					OrderId:   updatedOrder.ID,
					Status:    updatedOrder.Status,
					UpdatedAt: timestamppb.New(updatedOrder.CreatedAt), // Используем время из БД
				})

				if err != nil {
					log.Printf("Error sending to gRPC stream: %v", err)
					return err
				}
			}
		}
	}
}
