package gateway

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// r.GET("/blocks", getBlocks)
	r.POST("/playlist", createPlaylist)
}
