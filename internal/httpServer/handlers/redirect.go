package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectUrl(c *gin.Context) {
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

	c.Redirect(http.StatusMovedPermanently, url)
}
