// backend/grpc/server/paste_service.go
package server

import (
	"context"
	"gopaste/backend/models"
	pb "gopaste/proto"
	"gopaste/storage"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PasteServer struct {
	pb.UnimplementedPasteServiceServer
	store *storage.SQLiteStore
}

func NewPasteServer() *PasteServer {
	store, err := storage.NewSQLiteStore("./pastes.db")
	if err != nil {
		return nil
	}
	return &PasteServer{store: store}
}

func (s *PasteServer) CreatePaste(ctx context.Context, req *pb.CreatePasteRequest) (*pb.PasteResponse, error) {
	paste := &models.Paste{
		ID:        uuid.New().String(),
		Content:   req.Content,
		CreatedAt: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(req.ExpireHours) * time.Hour).Unix(),
	}

	if err := s.store.SavePaste(paste); err != nil {
		return nil, err
	}
	return &pb.PasteResponse{
		Id:        paste.ID,
		Content:   paste.Content,
		CreatedAt: paste.CreatedAt,
		ExpiresAt: paste.ExpiresAt,
	}, nil
}

func (s *PasteServer) GetPaste(ctx context.Context, req *pb.GetPasteRequest) (*pb.PasteResponse, error) {

	paste, err := s.store.GetPaste(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "paste not found")
	}
	if time.Now().Unix() > paste.ExpiresAt {
		return nil, status.Error(codes.NotFound, "paste expired")
	}
	return &pb.PasteResponse{
		Id:        paste.ID,
		Content:   paste.Content,
		CreatedAt: paste.CreatedAt,
		ExpiresAt: paste.ExpiresAt,
	}, nil
}
