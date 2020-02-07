package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vijaykrishnavanshi/go-grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Greet funxtion was invoked with %v\n", req)
	firstNum := req.GetInput().GetFirstNum()
	secondNum := req.GetInput().GetSecondNum()
	fmt.Printf("First Num: %v", firstNum)
	fmt.Printf("Second Num: %v", secondNum)
	result := firstNum + secondNum
	res := &calculatorpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Unable to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to serve: %v", err)
	}
}
