package main

import (
	"google.golang.org/grpc"
	"grpcadder/hello"
	"log"
	"net"
)

func main() {
	// listen to the network
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %+v", err)
	}

	// prepare the grpc server
	grpcServer := grpc.NewServer()
	s := hello.Server{} // handler

	// register the service server by passing in the grpc server
	hello.RegisterHelloServiceServer(grpcServer, &s)

	// call the serve function
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %+v", err)
	}
}
