package routes

import (
	"github.com/lxnewayfarer/shortlinks/handlers"
	"net/http"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", handlers.Ping)
	mux.HandleFunc("POST /shorten", handlers.Shorten)

	return mux
}