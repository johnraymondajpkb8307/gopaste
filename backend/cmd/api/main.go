// backend/cmd/api/main.go
package main

import (
	"gopaste/backend/api"
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

	// 设置路由
	r := api.SetupRouter(conn)

	// 启动服务器
	r.Run(":8080")
}
