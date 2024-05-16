package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	host := conf.Database.Host
	port := conf.Database.Port
	user := conf.Database.User
	password := conf.Database.Password
	database := conf.Database.Name

	// Connect to postgres
	dsn := fmt.Sprintf("host=%s port=%s, user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
}
