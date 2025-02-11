package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlDeleter interface {
	DeleteUrl(alias string) error
}

var urlDeleter UrlDeleter

func InitUrlDeleter(deleter UrlDeleter) {
	urlDeleter = deleter
}

func DeleteUrl(c *gin.Context) {
	alias := c.Param("alias")
	if alias == "" {
		c.JSON(http.StatusBadRequest, Response{Error: "Invalid request"})
		return
	}
	err := urlDeleter.DeleteUrl(alias)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"alias": alias})
}
