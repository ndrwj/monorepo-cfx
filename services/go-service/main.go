package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Log setiap request
	log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello from Go Service! 🚀")
}

func main() {
	http.HandleFunc("/", handler)

	// Log saat service mulai
	log.Println("Go service running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}