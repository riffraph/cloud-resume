package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type App struct {
	Router *http.ServeMux

	siteStatistics *SiteStatistics
}

func (a *App) Initialize() {
	log.Print("Initializing server...")

	a.siteStatistics = &SiteStatistics{Visits: 20}

	a.initialiseRoutes()
}

func (a *App) Run(addr string) {
	log.Printf("listening on port %s", addr)

	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		log.Fatal(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//
// ROUTE HANDLERS
//

func (a *App) getSiteStatistics(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, a.siteStatistics)
}

func (a *App) addVisit(w http.ResponseWriter, r *http.Request) {
	err := a.siteStatistics.addVisit()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, a.siteStatistics)
}

func (a *App) initialiseRoutes() {
	a.Router = http.NewServeMux()
	a.Router.HandleFunc("/visits", a.getSiteStatistics)
	a.Router.HandleFunc("/addvisit", a.addVisit)
}
