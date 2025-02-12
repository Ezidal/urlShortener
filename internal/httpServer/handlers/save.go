package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"time"

	mr "math/rand"

	"github.com/gin-gonic/gin"
)

type UrlSaver interface {
	SaveUrl(instUrl, alias string) error
}

var urlSaver UrlSaver

func InitUrlSaver(saver UrlSaver, logger *slog.Logger) {
	urlSaver = saver
	log = logger
}

func SaveUrl(c *gin.Context) {
	var req Request
	var err error
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: "Invalid request"})
		log.Error("Invalid request")
		return
	}

	if req.Url == "" {
		c.JSON(http.StatusBadRequest, Response{Error: "URL is required"})
		log.Error("URL is required")
		return
	}

	if req.Alias == "" {
		req.Alias, err = generateRandomAlias()
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
			log.Error(err.Error())
			return
		}
	}

	err = urlSaver.SaveUrl(req.Url, req.Alias)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		log.Error(err.Error())
		return
	}
	log.Info("Save alias: " + req.Alias + " with url: " + req.Url)
	c.JSON(http.StatusOK, Response{Alias: req.Alias})
}

func generateRandomAlias() (string, error) {
	mr.Seed(time.Now().UnixNano())
	length := mr.Intn(8) + 3
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}
