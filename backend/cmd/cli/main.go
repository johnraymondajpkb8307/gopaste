// backend/cmd/cli/main.go
package main

import (
	"flag"
	pb "gopasteb/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	content := flag.String("content", "", "Content to paste")
	flag.Parse()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPasteServiceClient(conn)
	// Implementation

}
