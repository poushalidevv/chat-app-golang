package main

import (
	"log"
	"net/http"
	"chat-app/config"
	"chat-app/internal/handlers"
)

func main() {
	cfg := config.LoadConfig()

	http.HandleFunc("/ws", handlers.WebSocketHandler)

	log.Printf("Server starting on %s...\n", cfg.ServerAddress)
	err := http.ListenAndServe(cfg.ServerAddress, nil)
	if err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}
}
