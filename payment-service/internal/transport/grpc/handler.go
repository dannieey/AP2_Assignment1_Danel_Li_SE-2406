package grpc

import (
	"context"
	"payment-service/internal/usecase"

	"github.com/dannieey/assignment2-generated/payment"
	"google.golang.org/protobuf/types/known/timestamppb" // Добавь это
)

type PaymentGRPCHandler struct {
	payment.UnimplementedPaymentServiceServer
	usecase *usecase.PaymentUseCase
}

func NewPaymentGRPCHandler(uc *usecase.PaymentUseCase) *PaymentGRPCHandler {
	return &PaymentGRPCHandler{
		usecase: uc,
	}
}

func (h *PaymentGRPCHandler) ProcessPayment(ctx context.Context, req *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	p, err := h.usecase.ProcessPayment(req.GetOrderId(), int64(req.GetAmount()))
	if err != nil {
		return &payment.PaymentResponse{
			Status:    "FAILED",
			CreatedAt: timestamppb.Now(),
		}, nil
	}

	return &payment.PaymentResponse{
		TransactionId: p.TransactionID,
		Status:        p.Status,
		CreatedAt:     timestamppb.Now(),
	}, nil
}
