package main

import (
	"log"
	"net/http"

	"go-web/internal/config"
	"go-web/internal/router"
)

func main() {
	cfg := config.Load()
	handler := router.New()

	log.Printf("server starting on %s", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, handler); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
