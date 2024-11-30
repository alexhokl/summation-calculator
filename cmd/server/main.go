package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/alexhokl/summation-calculator/internal"
	"github.com/alexhokl/summation-calculator/proto"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	proto.RegisterSummationServiceServer(grpcServer, &internal.Server{})

	log.Printf("Server is running on port %d\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
