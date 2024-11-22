// backend/grpc/server/paste_service.go
package server

import (
	"context"
	pb "gopaste/proto"
	"time"

	"github.com/google/uuid"
)

type PasteServer struct {
	pb.UnimplementedPasteServiceServer
	// 可以添加数据存储相关字段
}

func NewPasteServer() *PasteServer {
	return &PasteServer{}
}

func (s *PasteServer) CreatePaste(ctx context.Context, req *pb.CreatePasteRequest) (*pb.PasteResponse, error) {
	id := uuid.New().String()
	now := time.Now()
	expiresAt := now.Add(time.Duration(req.ExpireHours) * time.Hour)

	return &pb.PasteResponse{
		Id:        id,
		Content:   req.Content,
		CreatedAt: now.Unix(),
		ExpiresAt: expiresAt.Unix(),
	}, nil
}

func (s *PasteServer) GetPaste(ctx context.Context, req *pb.GetPasteRequest) (*pb.PasteResponse, error) {
	// 这里需要实现从存储中获取paste的逻辑
	// 暂时返回模拟数据
	return &pb.PasteResponse{
		Id:        req.Id,
		Content:   "Sample content",
		CreatedAt: time.Now().Unix(),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}, nil
}
