package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getClientIP(r *http.Request) string {
	// Ambil X-Forwarded-For, biasanya dipisah koma jika ada lebih dari satu proxy
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// ambil IP pertama (origin client)
		parts := strings.Split(xff, ",")
		return strings.TrimSpace(parts[0])
	}
	// fallback ke remote address
	return r.RemoteAddr
}

func handler(w http.ResponseWriter, r *http.Request) {
	clientIP := getClientIP(r)
	log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, clientIP)
	fmt.Fprintf(w, "Hello from Go Service! 🚀")
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Go service running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}