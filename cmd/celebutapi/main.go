package main

import (
	"celebut-api/configs"
	"celebut-api/internal/app"
	"log"
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
