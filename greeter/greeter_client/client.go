package main

import (
	"context"
	"fmt"
	"log"

	"../greetpb"

	"google.golang.org/grpc"
)

func main() {

	// Use withInsecure option to ignore SSL certs - creates insecure connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("failed to connect to server: ", err)
	}
	// Close connection once main is finished execution
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	res, err := createUnaryCall(c)

	fmt.Println(res)
}

func createUnaryCall(c greetpb.GreetServiceClient) (*greetpb.GreetResponse, error) {
	// Create request object - a greet request which holds a greeting
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName:       "Devon",
			LastName:        "Edwards Joseph",
			FavouriteCoffee: "Cappuccino",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("Greet request raised an error: ", err)
		return nil, err
	}
	log.Print("response from Greet: ", res)
	return res, nil
}
