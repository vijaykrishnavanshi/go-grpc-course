package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
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

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Server Streaming RPC")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vijay",
			LastName:  "Krishnavanshi",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while requesting server streaming: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while recieving message: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.Result)
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Server Streaming RPC")
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Vijay",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Vijay 1",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Vijay 2",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Vijay 4",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Vijay 5",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while requesting client streaming: %v", err)
	}
	for _, req := range requests {
		log.Printf("Sending Client Request: %v", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while recieving client streaming: %v", err)
	}
	log.Printf("Message Recieved: %v", msg.GetResult())
}
