package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/VincentBaron/youtube_to_tracklist/backend/gateway"
	"github.com/VincentBaron/youtube_to_tracklist/backend/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func LoadConfig(file string) (models.Config, error) {
// 	var config models.Config
// 	data, err := ioutil.ReadFile(file)
// 	if err != nil {
// 		return config, err
// 	}
// 	err = yaml.Unmarshal(data, &config)
// 	return config, err
// }

func main() {
	configsFile, err := os.ReadFile("configs.yml")
	if err != nil {
		// handle error
	}

	// Unmarshal the configsFile data into a Config struct
	var config models.Config
	err = yaml.Unmarshal(configsFile, &config)
	if err != nil {
		// handle error
	}
	log.Printf("config: %#v", config)

	// Extract the necessary information from the Config struct
	host := config.Database.Host
	port := config.Database.Port
	user := config.Database.User
	password := config.Database.Password
	database := config.Database.Name
	dsn := fmt.Sprintf("host=%s port=%s, user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting the DB(): %s", err)
	}
	defer sqlDB.Close()

	// db.AutoMigrate(models.Models())
	// db.AutoMigrate(&models.Block{})

	r := gin.Default()
	dbMiddleware := func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
	r.Use(dbMiddleware)

	// Set up the API endpoints
	gateway.SetupRoutes(r)

	// Start the server
	log.Printf("Server started at http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
