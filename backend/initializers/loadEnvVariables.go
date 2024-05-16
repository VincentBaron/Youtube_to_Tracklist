package initializers

import (
	"os"

	"github.com/VincentBaron/youtube_to_tracklist/backend/models"
	"gopkg.in/yaml.v2"
)

var Conf models.Config

func LoadEnvVariables() {
	configsFile, err := os.ReadFile("configs.yml")
	if err != nil {
		// handle error
	}

	// Unmarshal the configsFile data into a Config struct
	err = yaml.Unmarshal(configsFile, &Conf)
	if err != nil {
		// handle error
	}
}
