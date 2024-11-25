// backend/grpc/server/user_service.go
package server

import (
	"context"
	"gopaste/middleware"
	pb "gopaste/proto"
	"gopaste/storage"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	store *storage.SQLiteStore
}

func NewUserServer(store *storage.SQLiteStore) *UserServer {

	return &UserServer{store: store}
}

func (s *UserServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserResponse, error) {
	// Check if username exists
	if exists := s.store.UsernameExists(req.Username); exists {
		return nil, status.Error(codes.AlreadyExists, "username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	// Create user
	user, err := s.store.CreateUser(req.Username, string(hashedPassword))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &pb.UserResponse{
		Id:       user.ID,
		Username: user.Username,
	}, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.TokenResponse, error) {
	// Get user by username
	user, err := s.store.GetUserByUsername(req.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid password")
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate token")
	}

	return &pb.TokenResponse{
		Token: token,
	}, nil
}

func (s *UserServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	// Add token to blacklist if implementing token invalidation
	return &pb.LogoutResponse{
		Success: true,
	}, nil
}
