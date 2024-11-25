// backend/cmd/grpc/main.go
package main

import (
	"gopaste/backend/grpc/server"
	pb "gopaste/proto"
	"gopaste/storage"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	store, err := storage.NewSQLiteStore("./pastes.db")
	if err != nil {
		log.Fatalf("failed to create store: %v", err)
	}
	pb.RegisterPasteServiceServer(s, server.NewPasteServer(store))
	pb.RegisterUserServiceServer(s, server.NewUserServer(store))
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
