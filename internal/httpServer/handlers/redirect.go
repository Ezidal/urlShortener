package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectUrl(c *gin.Context) {
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
	log.Info("Redirect to url: " + url)
	c.Redirect(http.StatusMovedPermanently, url)
}
