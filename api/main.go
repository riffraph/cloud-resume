package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type SiteStatistics struct {
	Visits int `json:"Visits"`
}

var siteStats SiteStatistics

func main() {
	log.Print("starting server...")

	siteStats = SiteStatistics{Visits: 20}

	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/addvisit", addVisitHandler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	// add code to read from a database here
	json.NewEncoder(w).Encode(siteStats)
}

func addVisitHandler(w http.ResponseWriter, r *http.Request) {
	// add to the visit count
	siteStats.Visits++

	json.NewEncoder(w).Encode(siteStats)
}
