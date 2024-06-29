package initializers

import "github.com/VincentBaron/youtube_to_tracklist/backend/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Set{})
}
