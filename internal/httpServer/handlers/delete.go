package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlDeleter interface {
	DeleteUrl(alias string) error
}

var urlDeleter UrlDeleter

func InitUrlDeleter(deleter UrlDeleter, logger *slog.Logger) {
	urlDeleter = deleter
	log = logger
}

func DeleteUrl(c *gin.Context) {
	alias := c.Param("alias")
	if alias == "" {
		c.JSON(http.StatusBadRequest, Response{Error: "Invalid request"})
		log.Error("Invalid request")
		return
	}
	err := urlDeleter.DeleteUrl(alias)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{Error: err.Error()})
		log.Error(err.Error())
		return
	}
	log.Info("Deleted alias: " + alias)
	c.JSON(http.StatusOK, gin.H{"alias": alias})
}
