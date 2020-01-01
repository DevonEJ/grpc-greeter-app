package main

import (
	"context"
	"io"
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

	//res, err := createUnaryCall(c)

	createServerStreamingCall(c)

	//fmt.Println(res)
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

func createServerStreamingCall(c greetpb.GreetServiceClient) {
	// Create request object - a greet request which holds a greeting
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName:       "Devon",
			LastName:        "Edwards Joseph",
			FavouriteCoffee: "Cappuccino",
		},
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal("an error occured whilst calling server streaming GreetManyTimes: ", err)
	}

	// Loop over the result message stream:
	for {
		mssg, err := stream.Recv()

		if err == io.EOF {
			// End of message stream has been reached
			break
		}
		if err != nil {
			log.Fatal("error occurred whilst receiving stream: ", err)
		}

		log.Print("Response from GreetManyTimes: ", mssg)
	}

}
