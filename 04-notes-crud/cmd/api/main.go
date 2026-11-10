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
	// step 1: load our env configurations
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// step 2: use those configs to connect to our DB
	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// step 2.2: also run disconnect as deferred func which runs after main finishes
	defer func() {
		if err := db.Disconnect(client); err != nil {
			log.Printf("failed to disconnect from database: %v", err)
		}
	}()

	// step 3: the router runs the app/server on given port and serves the requests per the routes
	router := server.NewRouter(database)

	portAddress := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("listening on %s", portAddress)
	if err := router.Run(portAddress); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
