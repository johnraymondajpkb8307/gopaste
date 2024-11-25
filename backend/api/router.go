// backend/api/router.go
package api

import (
	"gopaste/backend/api/handlers"
	"gopaste/middleware"
	pb "gopaste/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func SetupRouter(conn *grpc.ClientConn) *gin.Engine {
	r := gin.Default()
	pasteClient := pb.NewPasteServiceClient(conn)
	authClient := pb.NewUserServiceClient(conn)
	authHandler := handlers.NewAuthHandler(authClient)
	r.POST("/api/login", authHandler.Login)
	r.POST("/api/register", authHandler.Register)
	auth := r.Group("/api")
	pasteHandler := handlers.NewPasteHandler(pasteClient)
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/api/paste", pasteHandler.CreatePaste)
		auth.GET("/api/paste/:id", pasteHandler.GetPaste)
	}
	return r
}
