package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlGetter interface {
	GetUrl(alias string) (string, error)
	GetAllUrls() (map[string]string, error)
}

var urlGetter UrlGetter

func InitUrlGetter(getter UrlGetter) {
	urlGetter = getter
}

func GetUrl(c *gin.Context) {
	alias := c.Param("alias")
	if alias == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Alias is required"})
		return
	}

	url, err := urlGetter.GetUrl(alias)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}

func GetAllUrls(c *gin.Context) {
	urls, err := urlGetter.GetAllUrls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, urls)
}
