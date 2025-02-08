package main

import (
	"log"

	"github.com/sawmeraw/ims/internal/server"
)

func main() {
	s, err := server.Init()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	s.SetupMiddleware().SetupRoutes()

	if err := s.Start(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
