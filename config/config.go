package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Server  Server
	MongoDB MongoDB
	Redis   Redis
	Jwt     Jwt
}

type MongoDB struct {
	ConnectionUrl string
}

type Server struct {
	Port string
}

type Jwt struct {
	Secret    string
	Expires   string
	Algorithm string
}

type Redis struct {
	ConnectionUrl string
}

const envPath = ".env"

func New() Config {
	var config Config
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	config.MongoDB = MongoDB{
		ConnectionUrl: os.Getenv("MONGO_CONNECTION_URL"),
	}

	config.Server = Server{
		Port: os.Getenv("SERVER_PORT"),
	}

	config.Redis = Redis{
		ConnectionUrl: os.Getenv("REDIS_CONNECTION_URL"),
	}

	config.Jwt = Jwt{
		Algorithm: os.Getenv("JWT_ALGORITHM"),
		Expires:   os.Getenv("JWT_EXPIRES"),
		Secret:    os.Getenv("JWT_SECRET"),
	}
	return config
}

func (c *Config) Print() {
	fmt.Println(*c)
}
