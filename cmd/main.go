package main

import (
	"log"
	"os"
	"urlShortener/internal/config"
	"urlShortener/internal/logger"
	"urlShortener/internal/logger/er"
	"urlShortener/internal/storage/sqlite"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log := logger.SetLogger(config.Environment)

	log.Info("Service starting")

	storage, err := sqlite.New(config.Storage_path)
	if err != nil {
		log.Error("failed to init db", er.Err(err))
		os.Exit(1)
	}
	// err = storage.SaveUrl("https://google.com", "gogle")
	storage.SaveUrl("popopoppo", "pop")
	err = storage.DeleteUrl("pop")
	if err != nil {
		log.Error("failed to delete db", er.Err(err))
	} else {
		log.Info("успех")
	}
}
