package main

import (
	"log"
	"net"

	"../greetpb"

	"google.golang.org/grpc"
)

type Server struct{}

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

	log.Print("successfully listening on port: ", port)

}
