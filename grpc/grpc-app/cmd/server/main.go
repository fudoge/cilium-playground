package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	testv1 "example.com/test/v1"
)

type server struct {
	testv1.UnimplementedTestServiceServer
}

func (s *server) Ping(_ context.Context, req *testv1.PingRequest) (*testv1.PingResponse, error) {
	return &testv1.PingResponse{
		Message:       req.GetMessage(),
		ServerTimeMs: time.Now().UnixMilli(),
	}, nil
}

func main() {
	addr := flag.String("addr", ":50051", "listen address")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	testv1.RegisterTestServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	log.Printf("grpc server listening on %s", *addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
