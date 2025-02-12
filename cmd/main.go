package main

import (
	"log"
	"os"
	"urlShortener/internal/config"
	"urlShortener/internal/httpServer/handlers"
	"urlShortener/internal/logger"
	"urlShortener/internal/logger/er"
	"urlShortener/internal/storage/sqlite"

	"github.com/gin-gonic/gin"
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

	r := gin.New()
	// r.Use(sloggin.New(log))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	handlers.InitUrlSaver(storage)
	handlers.InitUrlGetter(storage)
	handlers.InitUrlDeleter(storage)

	r.POST("/save", handlers.SaveUrl)
	r.GET("/get/:alias", handlers.GetUrl)
	r.GET("/get/all", handlers.GetAllUrls)
	r.DELETE("/delete/:alias", handlers.DeleteUrl)
	r.GET("/:alias", handlers.RedirectUrl)
	r.Run(":8080")
}
