package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"../greetpb"

	"google.golang.org/grpc"
)

//Server implements the GreetServiceServer interface from the pb.go file, with the below methods
type Server struct{}

//Greet implements a unary api to send a single greeting response
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

//GreetManyTimes implements a streaming api to send many greeting responses
func (*Server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Print("Request made to GreetManyTimes function: ", req)
	// Extract fields from the protobuf messages in the client request
	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		response := "Hello " + firstName + ". You are number: " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Response: response,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
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

	log.Print("listening on port: ", port)

	if err := server.Serve(listener); err != nil {
		log.Fatal("Failed to connect: ", err)
	}

}
