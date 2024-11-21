// backend/api/router.go
package api

import (
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    
    r.POST("/api/paste", createPaste)
    r.GET("/api/paste/:id", getPaste)
    
    return r
}