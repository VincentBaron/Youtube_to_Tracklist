package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/VincentBaron/youtube_to_tracklist/backend/gateway"
	"github.com/VincentBaron/youtube_to_tracklist/backend/initializers"
	"github.com/VincentBaron/youtube_to_tracklist/backend/middlewares"
	"github.com/VincentBaron/youtube_to_tracklist/backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func LoadConfig(file string) (models.Config, error) {
	var config models.Config
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func main() {

	// Set up the Gin router
	r := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// Set up the API endpoints
	// r.GET("/blocks", getBlocks)
	r.POST("/playlist", middlewares.RequireAuth, gateway.CreatePlaylist)
	r.GET("/validate", middlewares.RequireAuth, gateway.Validate)
	// r.GET("/spotify-login", handler.spotifyLoginHandler)
	// r.GET("/home", handler.homeHandler)
	// r.GET("/callback", handler.handleCallback)
	r.POST("/signup", gateway.Signup)
	r.POST("/login", gateway.Login)
	r.GET("/callback", gateway.CallbackHandler)
	// r.GET("/status", handler.handleStatus)
	// r.POST("/store-token", storeTokenHandler)
	r.Run(":8080")

	// Start the server
	log.Printf("Server started at http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
