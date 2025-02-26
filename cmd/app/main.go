package main

import (
	"log"

	"Go-Boilerplate/internal/config"
	"Go-Boilerplate/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from a specific .env file.
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found at ../../.env")
	}

	// Load application configuration.
	cfg := config.LoadConfig()

	// Initialize and start the server.
	srv := server.NewServer(cfg)
	if err := srv.Start(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
