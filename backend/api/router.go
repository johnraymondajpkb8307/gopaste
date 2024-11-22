// backend/api/router.go
package api

import (
	"gopaste/backend/api/handlers"
	pb "gopaste/proto"

	"github.com/gin-gonic/gin"
)

func SetupRouter(grpcClient pb.PasteServiceClient) *gin.Engine {
	r := gin.Default()
	pasteHandler := handlers.NewPasteHandler(grpcClient)
	r.POST("/api/paste", pasteHandler.CreatePaste)
	r.GET("/api/paste/:id", pasteHandler.GetPaste)

	return r
}
