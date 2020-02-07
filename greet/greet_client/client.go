package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vijaykrishnavanshi/go-grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to Dial: %v", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vijay",
			LastName:  "Krishnavanshi",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Response from greet: %v", res.Result)
}
