package main

import (
	"log"
	"net/http"

	"auth-lab-go/internal/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/auth/oidc/login", handlers.OIDCLoginHandler)
	http.HandleFunc("/auth/oidc/callback", handlers.OIDCCallbackHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
