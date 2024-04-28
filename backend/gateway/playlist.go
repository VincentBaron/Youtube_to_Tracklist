package gateway

import (
	"net/http"
	"net/url"

	"github.com/VincentBaron/youtube_to_tracklist/backend/dto"
	"github.com/gin-gonic/gin"
)

func createPlaylist(c *gin.Context) {
	var req dto.PostPlaylistReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := url.Parse(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}
	queryParams := u.Query()
	videoID := queryParams.Get("v")

	// youtubeService, err := youtube.NewService(c)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// video := youtubeService.Videos.List([]string{"snippet"}).Id(videoID)

	// db, _ := c.Get("DB")
	// blockService := blocks.Block{DB: db.(*gorm.DB)}
	// resp, err := blockService.PostPlaylist(req.Block)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, videoID)
}
