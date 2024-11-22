// backend/cmd/test/grpc_test.go
package main

import (
	"context"
	pb "gopaste/proto"
	"log"
	"testing"

	"google.golang.org/grpc"
)

func testGRPC(_ *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPasteServiceClient(conn)

	// 测试创建paste
	createResp, err := client.CreatePaste(context.Background(), &pb.CreatePasteRequest{
		Content:     "Test content",
		ExpireHours: 24,
	})
	if err != nil {
		log.Fatalf("CreatePaste failed: %v", err)
	}
	log.Printf("Created paste: %v", createResp)

	// 测试获取paste
	getResp, err := client.GetPaste(context.Background(), &pb.GetPasteRequest{
		Id: createResp.Id,
	})
	if err != nil {
		log.Fatalf("GetPaste failed: %v", err)
	}
	log.Printf("Retrieved paste: %v", getResp)
}
