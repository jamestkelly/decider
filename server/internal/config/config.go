package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port int
}

func New() *Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Invalid PORT environment variable")
	}

	return &Config{
		Port: port,
	}
}
