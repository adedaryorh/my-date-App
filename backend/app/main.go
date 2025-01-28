package main

import (
	"log"

	"backend.app/configs"
	"backend.app/internal/app"
)

func init() {
}

func main() {
	// Configuration
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
