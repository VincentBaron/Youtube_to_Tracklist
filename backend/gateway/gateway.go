package gateway

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// r.GET("/blocks", getBlocks)
	r.POST("/playlist", createPlaylist)
	r.GET("/spotify-login", spotifyLoginHandler)
	r.GET("/home", homeHandler)
	r.GET("/callback", callbackHandler)
	// r.POST("/store-token", storeTokenHandler)
}
