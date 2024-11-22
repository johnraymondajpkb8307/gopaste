package handlers

import (
	pb "gopaste/proto"

	"github.com/gin-gonic/gin"
)

type PasteHandler struct {
	grpcClient pb.PasteServiceClient
}

func NewPasteHandler(
	grpcClient pb.PasteServiceClient,
) *PasteHandler {
	return &PasteHandler{
		grpcClient: grpcClient,
	}
}

func (h *PasteHandler) CreatePaste(c *gin.Context) {
	// Implementation
	var req pb.CreatePasteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := h.grpcClient.CreatePaste(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

func (h *PasteHandler) GetPaste(c *gin.Context) {
	id := c.Param("id")
	req := &pb.GetPasteRequest{
		Id: id,
	}
	resp, err := h.grpcClient.GetPaste(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
