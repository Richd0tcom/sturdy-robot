package config

import (
	"fmt"
	// "log"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
	DbDriver string `mapstructure:"DB_DRIVER"`
	DbUri string `mapstructure:"DB_URI"`
	ServerAddress string `mapstructure:"SERVER_URL"`
	Enviroment string `mapstructure:"GO_ENV"`
	SslCert string `mapstructure:"SSL_CERT"`
}

func LoadConfig(path string) (*Config,  error) {
	err:= godotenv.Load(path)
  if err != nil {
    fmt.Println("Error loading .env file")
	
  }

   	dbDriver:= os.Getenv("DB_DRIVER")
  	dbUri := os.Getenv("DB_URI")
  	address := os.Getenv("SERVER_URL")

	  enviroment:= os.Getenv("GO_ENV")

	  if enviroment== "production" {
		dbUri += fmt.Sprintf("&sslrootcert=%s", os.Getenv("SSL_CERT"))
	  }
  

	return &Config{
		DbDriver: dbDriver,
		DbUri: dbUri,
		ServerAddress: address,
		Enviroment: os.Getenv("GO_ENV"),
	}, nil
}