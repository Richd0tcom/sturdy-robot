package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	DbDriver string `mapstructure:"DB_DRIVER"`
	DbUri string `mapstructure:"DB_URI"`
	ServerAddress string `mapstructure:"SERVER_URL"`
}

func LoadConfig(path string) (*Config,  error) {
	err:= godotenv.Load(path)
  if err != nil {
    log.Fatal("Error loading .env file")
	
  }

   	dbDriver:= os.Getenv("DB_DRIVER")
  	dbUri := os.Getenv("DB_URI")
  	address := os.Getenv("SERVER_URL")


	return &Config{
		DbDriver: dbDriver,
		DbUri: dbUri,
		ServerAddress: address,
	}, nil
}