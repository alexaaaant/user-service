package config

import (
	"log"
	"os"
)

type Config struct {
	DBURL string
}

func Load() Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	return Config{
		DBURL: dbURL,
	}
}
