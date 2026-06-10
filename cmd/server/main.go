package main

import (
	"log"

	"go-web/internal/config"
	"go-web/internal/router"
)

func main() {
	cfg := config.Load()
	engine := router.New()

	log.Printf("server starting on %s", cfg.Addr)
	if err := engine.Run(cfg.Addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
