package handlers

import (
	"net/http"
	"log/slog"
	"encoding/json"
	"github.com/lxnewayfarer/shortlinks/lib"
)

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	slog.Info("Ping")
	JSONResponse(w, http.StatusOK, map[string]string{
		"response": "pong",
	})
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

	link := r.FormValue("link")

	shortenLink, err := lib.ShortenLink(link)
	if err != nil {
		http.Error(w, "Can not process link", http.StatusBadRequest)
        return
	}

	JSONResponse(w, http.StatusOK, map[string]string{
		"response": shortenLink,
	})
}