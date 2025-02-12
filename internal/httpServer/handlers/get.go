package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlGetter interface {
	GetUrl(alias string) (string, error)
	GetAllUrls() (map[string]string, error)
}

var urlGetter UrlGetter

func InitUrlGetter(getter UrlGetter, logger *slog.Logger) {
	urlGetter = getter
	log = logger
}

func GetUrl(c *gin.Context) {
	alias := c.Param("alias")
	if alias == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Alias is required"})
		log.Error("Alias is required")
		return
	}

	url, err := urlGetter.GetUrl(alias)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		log.Error(err.Error())
		return
	}
	log.Info("Getted alias: " + alias)
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func GetAllUrls(c *gin.Context) {
	urls, err := urlGetter.GetAllUrls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Error(err.Error())
		return
	}
	log.Info("Getted all urls")
	c.JSON(http.StatusOK, urls)
}
