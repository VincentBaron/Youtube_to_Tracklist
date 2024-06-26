package gateway

import (
	"net/http"

	"github.com/VincentBaron/youtube_to_tracklist/backend/initializers"
	"github.com/VincentBaron/youtube_to_tracklist/backend/models"
	"github.com/gin-gonic/gin"
)

func GetSets(c *gin.Context) {

	// Create a new slice to store the playlist names
	var sets []models.Set

	initializers.DB.Find(&sets)

	// Return the list of playlist names
	c.JSON(http.StatusOK, gin.H{"sets": sets})
}
