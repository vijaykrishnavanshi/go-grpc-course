package main

import (
	"context"
	"fmt"
	"log"

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
	doUnary(c)
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
