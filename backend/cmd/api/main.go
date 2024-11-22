// backend/cmd/api/main.go
package main

import (
	"gopaste/backend/api"
	pb "gopaste/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// 连接gRPC服务
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 创建gRPC客户端
	grpcClient := pb.NewPasteServiceClient(conn)

	// 设置路由
	r := api.SetupRouter(grpcClient)

	// 启动服务器
	r.Run(":8080")
}
