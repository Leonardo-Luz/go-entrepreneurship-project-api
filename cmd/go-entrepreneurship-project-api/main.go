package main

import (
	"fmt"
	"log"

	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/config"
	"github.com/leonardo-luz/go-entrepreneurship-project-api/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro loading .env file, %v", err)
	}

	db, err := config.ConnectDB(cfg)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := config.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database models: %v", err)
	}

	router, err := router.SetupRouter(db, cfg)

	if err != nil {
		log.Fatalf("Failed to setup router: %v", err)
	}

	addr := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
