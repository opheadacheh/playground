package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	helloworldpb "project/proto"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := helloworldpb.NewGreeterClient(conn)
	ctx := context.Background()

	// how to perform healthcheck request manually:
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	resp, err := c.SayHello(ctx, &helloworldpb.HelloRequest{Name: "A"})
	if err != nil {
		log.Fatalf("SayHello failed %+v", err)
	}

	log.Printf("response: %v\n", resp)
}
