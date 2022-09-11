package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TranQuocToan1996/goWebSocket/internal/handlers"
)

func main() {
	routes := newMux()
	port := 8080

	go handlers.ListenToWsChan()

	log.Printf("starting server on port %v!", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), routes)
	// http.ListenAndServeTLS()
}
