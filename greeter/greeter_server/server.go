package main

import (
	"context"
	"log"
	"net"

	"../greetpb"

	"google.golang.org/grpc"
)

type Server struct{}

//Greet implements the GreetServiceServer interface (with a Greet method) from the pb.go file
func (*Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Print("Request made to Greet function: ", req)
	// Extract fields from the protobuf messages in the client request
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	coffee := req.GetGreeting().GetFavouriteCoffee()

	// Create response message
	response := "Hello " + firstName + " " + lastName + "! Have a " + coffee + " on me :)"
	resMssg := &greetpb.GreetResponse{
		Response: response,
	}
	return resMssg, nil
}

func main() {

	port := "0.0.0.0:50051"

	// Listen on default port for grpc
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Failed to listen on port: ", err)
	}

	server := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(server, &Server{})

	log.Print("attempting to listen on port: ", port)

	if err := server.Serve(listener); err != nil {
		log.Fatal("Failed to connect: ", err)
	}

}
