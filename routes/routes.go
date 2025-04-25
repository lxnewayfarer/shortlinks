package routes

import (
	"github.com/lxnewayfarer/shortlinks/handlers"
	"net/http"
)

func Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.Ping)

	return mux
}