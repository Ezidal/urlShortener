package handlers

import "log/slog"

type Request struct {
	Url   string `json:"url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Alias string `json:"alias"`
}

var log *slog.Logger
