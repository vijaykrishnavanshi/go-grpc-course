package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/vijaykrishnavanshi/go-grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to Dial: %v", err)
	}
	fmt.Printf("Connection Created: %v", cc)
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		Input: &calculatorpb.SumInput{
			FirstNum:  1,
			SecondNum: 1,
		},
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Unable to call: %v", err)
	}
	log.Printf("res: %v", res)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeDecompositionRequest{
		Number: 120,
	}
	resStream, err := c.PrimeDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Unable to call: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			log.Printf("Server stopped streaming")
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming message: %v", err)
		}
		log.Printf("Streamed Value: %v", msg.GetFactor())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	numbers := []int64{120, 233, 343, 4344, 3443}
	resStream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Unable to call: %v", err)
	}
	for _, number := range numbers {
		log.Printf("Sending Client Request: %v", number)
		resStream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
		time.Sleep(1000 * time.Millisecond)
	}
	msg, err := resStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Unable to call: %v", err)
	}
	log.Printf("Average: %v", msg.Average)
}
