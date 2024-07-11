package entities

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName          string
	Environment      string
	AppPort          string
	DBUri            string
	DBName           string
	APIVersion       string
	RestaurantApiUrl string
}

func NewConfig() (*Config, error) {
	loadEnvironment()

	config := &Config{
		AppName:          strings.TrimRight(os.Getenv("APP_NAME"), "\n\r"),
		Environment:      strings.TrimRight(os.Getenv("APP_ENV"), "\n\r"),
		AppPort:          strings.TrimRight(os.Getenv("APP_PORT"), "\n\r"),
		DBUri:            strings.TrimRight(os.Getenv("DB_URI"), "\n\r"),
		DBName:           strings.TrimRight(os.Getenv("DB_NAME"), "\n\r"),
		APIVersion:       strings.TrimRight(os.Getenv("API_VERSION"), "\n\r"),
		RestaurantApiUrl: strings.TrimRight(os.Getenv("RESTAURANT_API_URL"), "\n\r"),
	}

	printConfig(config)
	return config, nil
}

func loadEnvironment() {
	if os.Getenv("APP_ENV") == "production" {
		// Load .env file
		err := godotenv.Load()
		if err != nil {
			slog.Info("loading .env file not found")
		}
	} else if os.Getenv("APP_ENV") == "development" {
		// Load .env.development file
		err := godotenv.Load(".env.development")
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	} else {
		_ = os.Setenv("APP_NAME", "payment-api")
		_ = os.Setenv("APP_ENV", "default")
		_ = os.Setenv("APP_PORT", ":8010")
		_ = os.Setenv("DB_URI", "mongodb://<USER>:<PASSWORD>@localhost:27017/")
		_ = os.Setenv("DB_NAME", "tech_challenge_payment_db")
		_ = os.Setenv("API_VERSION", "4.0")
		_ = os.Setenv("RESTAURANT_API_URL", "http://localhost:8080")
	}

}

func printConfig(config *Config) {
	fmt.Println("*** Environments ***")
	fmt.Printf("App Name: %s\n", config.AppName)
	fmt.Printf("Environment: %s\n", config.Environment)
	fmt.Printf("App Port: %s\n", config.AppPort)
	fmt.Printf("DB DBUri: %s\n", config.DBUri)
	fmt.Printf("DB Name: %s\n", config.DBName)
	fmt.Printf("API version: %s\n", config.APIVersion)
	fmt.Printf("Restaurant Api URL: %s\n", config.RestaurantApiUrl)
}
