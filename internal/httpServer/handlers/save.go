package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	mr "math/rand"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Url   string `json:"url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Alias string `json:"alias"`
}

type UrlSaver interface {
	SaveUrl(instUrl, alias string) error
}

var urlSaver UrlSaver

func InitUrlSaver(saver UrlSaver) {
	urlSaver = saver
}

func SaveUrl(c *gin.Context) {
	var req Request
	var err error
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: "Invalid request"})
		return
	}

	if req.Url == "" {
		c.JSON(http.StatusBadRequest, Response{Error: "URL is required"})
		return
	}

	if req.Alias == "" {
		req.Alias, _ = generateRandomAlias()
	}

	err = urlSaver.SaveUrl(req.Url, req.Alias)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error()})
		return
	}

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
