package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConnection is the connection string to the database
	ConnectionString = ""
	// Port is the port where the database is running
	Port = 0
	// SecretKey is the key used to sign the JWT (json web token)
	SecretKey []byte
)

// Load loads the configuration
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		Port = 9000
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
