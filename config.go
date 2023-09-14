package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value from the variables key
func Config(key string) string {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
