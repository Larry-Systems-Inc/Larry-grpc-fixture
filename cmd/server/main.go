package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	fixturev1 "github.com/Larry-Systems-Inc/larry-grpc-fixture/gen/fixturev1"
	"github.com/Larry-Systems-Inc/larry-grpc-fixture/internal/server"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	fixturev1.RegisterFixtureServiceServer(
		grpcServer,
		&server.FixtureServer{},
	)

	reflection.Register(grpcServer)

	log.Println("Larry gRPC fixture listening on 127.0.0.1:50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
