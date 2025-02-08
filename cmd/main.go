package main

import (
	"log"
	"urlShortener/internal/config"
	"urlShortener/internal/logger"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log := logger.SetLogger(config.Environment)

}
