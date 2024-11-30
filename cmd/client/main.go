package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/alexhokl/summation-calculator/proto"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("Failed to run: %v", err)
		os.Exit(-1)
	}
}

func run() error {
	ctx := context.Background()

	if (len(os.Args) != 3) {
		return fmt.Errorf("Usage: %s <number1> <number2>", os.Args[0])
	}

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return fmt.Errorf("Failed to convert %s to integer: %v", os.Args[1], err)
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return fmt.Errorf("Failed to convert %s to integer: %v", os.Args[2], err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewSummationServiceClient(conn)

	r, err := c.Sum(ctx, &proto.SummationRequest{A: int64(a), B: int64(b)})
	if err != nil {
		return fmt.Errorf("Failed to complete API call: %v", err)
	}

	fmt.Printf("The sum is: %d\n", r.GetResult())

	return nil
}
