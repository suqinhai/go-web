package router

import (
	"net/http"

	"go-web/internal/api"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", api.Index)
	mux.HandleFunc("GET /health", api.Health)
	return mux
}
