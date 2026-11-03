package main

import (
	"04-notes-crud/internal/config"
	"04-notes-crud/internal/db"
	"04-notes-crud/internal/server"
	"fmt"
	"log"
)

// config -> db -> router -> run the server

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client, _, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer func() {
		if err := db.Disconnect(client); err != nil {
			log.Printf("failed to disconnect from database: %v", err)
		}
	}()

	router := server.NewRouter()

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("listening on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
