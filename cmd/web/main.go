package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes := newMux()
	port := 8080

	log.Println("starting server!")

	http.ListenAndServe(fmt.Sprintf(":%v", port), routes)
	// http.ListenAndServeTLS()
}
