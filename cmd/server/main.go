package main

import (
	"log"

	"go-web/internal/config"
	"go-web/internal/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	engine, err := router.New(cfg)
	if err != nil {
		log.Fatalf("init router failed: %v", err)
	}

	log.Printf("server starting on %s", cfg.Server.Addr)
	if err := engine.Run(cfg.Server.Addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
