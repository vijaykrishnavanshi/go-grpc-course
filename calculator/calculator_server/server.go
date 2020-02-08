package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

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

func (*server) PrimeDecomposition(req *calculatorpb.PrimeDecompositionRequest, stream calculatorpb.CalculatorService_PrimeDecompositionServer) error {
	fmt.Printf("Greet funxtion was invoked with %v\n", req)
	number := req.GetNumber()
	fmt.Printf("Number: %v\n", number)
	var k int64
	k = 2
	for number > 1 {
		if number%k == 0 {
			number = number / k // divide N by k so that we have the rest of the number left.
			res := &calculatorpb.PrimeDecompositionResponse{
				Factor: k,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
		} else {
			k = k + 1
		}

	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("Greet funxtion was invoked with client stream\n")
	sum := int64(0)
	numbersCount := int64(0)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			average := sum / numbersCount
			log.Printf("Average: %v", average)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("Error in client streaming: %v", err)
		}
		sum += msg.GetNumber()
		numbersCount++
	}
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
