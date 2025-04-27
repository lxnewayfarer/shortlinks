package routes

import (
	"net/http"

	"github.com/lxnewayfarer/shortlinks/handlers"
	"github.com/lxnewayfarer/shortlinks/storage"
)

func Init(rdb storage.RedisClient) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", handlers.Ping(rdb))
	mux.HandleFunc("POST /shorten", handlers.Shorten(rdb))

	return mux
}