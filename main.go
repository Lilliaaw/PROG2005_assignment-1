package main

import (
	"assignment-1/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	handler.SetStartTime() //sets time in which the programs starts

	// Handle port assignment (either based on environment variable, or local override)
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Set up handler endpoints
	http.HandleFunc(handler.DEFAULT_PATH, handler.EmptyHandler)
	http.HandleFunc(handler.UNIINFO_PATH, handler.InfoHandler)
	http.HandleFunc(handler.NEIGHBOURUNIS_PATH, handler.NeighbourHandler)
	http.HandleFunc(handler.DIAG_PATH, handler.DiagHandler)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
