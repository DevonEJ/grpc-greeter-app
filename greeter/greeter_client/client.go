package main

import (
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

	// TODO - Implement client requests
	fmt.Printf("created client: ", c)

}
