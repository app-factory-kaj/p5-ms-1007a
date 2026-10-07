package main

import (
	"log"
	"net/http"
	"os"

	"greeter/internal/httpapi"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: httpapi.NewRouter(),
	}

	log.Printf("greeter listening on :%s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
