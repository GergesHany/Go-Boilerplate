package config

import (
	"log"
	"os"
)

type Config struct {
	Port   string
	DBUser string
	DBPass string
	DBName string
	DBHost string
	DBPort string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("PORT environment variable not set, defaulting to 8080")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	return &Config{
		Port:   port,
		DBUser: dbUser,
		DBPass: dbPass,
		DBName: dbName,
		DBHost: dbHost,
		DBPort: dbPort,
	}
}
