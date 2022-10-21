package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcadder/hello"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %+v", err)
	}

	c := hello.NewHelloServiceClient(conn)
	m, err := c.Talkie(context.Background(), &hello.Message{Content: "Hello i am a client"})
	if err != nil {
		log.Fatalf("Failed to Talkie: %+v", err)
	}
	log.Print(m.Content)
}
