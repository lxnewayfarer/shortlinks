package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/lxnewayfarer/shortlinks/lib"
	"github.com/lxnewayfarer/shortlinks/storage"
)

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Redirect(rdb storage.RedisClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("path")
		link, err := rdb.Get(r.Context(), path).Result()
		if err != nil {
			JSONResponse(w, http.StatusNotFound, map[string]string{
				"response": "Not found",
			})
			return
		}
		http.Redirect(w, r, link, http.StatusSeeOther)
	}
}

func Ping(rdb storage.RedisClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Ping")
		JSONResponse(w, http.StatusOK, map[string]string{
			"response": "pong",
		})
	}
}

func Shorten(rdb storage.RedisClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		link := r.PostFormValue("link")

		shortenLink, err := lib.ShortenLink(r.Context(), rdb, link, lib.RandomInstance{})
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Can not process link", http.StatusBadRequest)
			return
		}

		JSONResponse(w, http.StatusOK, map[string]string{
			"response": shortenLink,
		})
	}
}
