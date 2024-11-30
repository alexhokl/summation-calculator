package internal

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexhokl/summation-calculator/proto"
)

type Server struct {
	proto.UnimplementedSummationServiceServer
}

func (s *Server) Sum(_ context.Context, req *proto.SummationRequest) (*proto.SummationResponse, error) {

	fmt.Printf("Received request: %v\n", req)

	a := req.GetA()
	b := req.GetB()

	if a < 0 || b < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid argument: a=%d, b=%d", a, b),
		)
	}

	sum := a + b

	fmt.Printf("Summation result: %d\n", sum)

   return &proto.SummationResponse{
	Result: sum,
   }, nil
}
