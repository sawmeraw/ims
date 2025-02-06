package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sawmeraw/ims/handlers"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.HandleFunc("/", handlers.Home)

	log.Println("Server is starting on: 8080")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatalf("Server failed to start: %s", err.Error())
	}

}
