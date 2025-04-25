package handlers

import (
	"net/http"
	"log/slog"
	"encoding/json"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	slog.Info("Ping")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]string{
		"response": "pong",
	})
}