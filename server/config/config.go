package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var Cfg Config

type Config struct {
	Port       string
	Host       string
	Mode       string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	AppUrl     string
}

func init() {
	loadConfig()
}

func loadConfig() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(filepath.Join(workingDir, ".env"))

	if err != nil {
		log.Fatal("Error loading .env file, using system environment variables")
	}

	Cfg = Config{
		Port:       os.Getenv("PORT"),
		Host:       os.Getenv("HOST"),
		Mode:       os.Getenv("MODE"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		AppUrl:     os.Getenv("APP_URL"),
	}
}
