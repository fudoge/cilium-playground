package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	testv1 "example.com/test/v1"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:50051", "server address")
	message := flag.String("message", "hello", "ping message")
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()

	client := testv1.NewTestServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.Ping(ctx, &testv1.PingRequest{Message: *message})
	if err != nil {
		log.Fatalf("ping: %v", err)
	}

	log.Printf("response message=%q server_time_ms=%d", resp.GetMessage(), resp.GetServerTimeMs())
}
