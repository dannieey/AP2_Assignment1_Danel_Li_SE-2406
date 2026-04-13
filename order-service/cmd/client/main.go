package main

import (
	"context"
	"io"
	"log"

	"github.com/dannieey/assignment2-generated/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := order.NewOrderServiceClient(conn)

	req := &order.OrderSubscriptionRequest{OrderId: ""}

	stream, err := client.SubscribeToOrderUpdates(context.Background(), req)
	if err != nil {
		log.Fatalf("error on subscribe: %v", err)
	}

	log.Println("--- Waiting for order updates ---")

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("cannot receive: %v", err)
		}
		log.Printf("UPDATE RECEIVED: Order ID: %s, Status: %s, Time: %s",
			resp.OrderId, resp.Status, resp.UpdatedAt.AsTime())
	}
}
