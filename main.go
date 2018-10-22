package main

import (
	"github.com/gorilla/mux"
	. "IMT2681-assignement-2/api"
	"IMT2681-assignement-2/mongodb"
	"log"
	"net/http"
	"os"
)


func main() {
	// Get port for Heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}


	//INIT DATABASE

	mongodb.Global = &mongodb.TracksMongoDB{
		"mongodb://heihade:heihade123@ds131983.mlab.com:31983/josteivhdb",
		"josteivhdb",
		"track",
	}

	mongodb.Global.Init()

	// Set up handlers
	r := mux.NewRouter()
	r.HandleFunc("/igcinfo/api", InfoHandler).Methods("GET")
	r.HandleFunc("/igcinfo/api/igc", GetAllId).Methods("GET")
	r.HandleFunc("/igcinfo/api/igc", AddTrack).Methods("POST")
	r.HandleFunc("/igcinfo/api/igc/{id:[0-9]+}", GetTrack).Methods("GET")
	r.HandleFunc("/igcinfo/api/igc/{id:[0-9]+}/{prop:[a-z_H]+}", GetTrackProp).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, r))
}