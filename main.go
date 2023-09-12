package main

import (
	"log"
	"net/http"
	"url_shortener/db"
	"url_shortener/handlers"

	"github.com/gorilla/mux"
)

func main() {
	const PORT = ":4000"
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/shortener", h.Shortener).Methods("POST")
	router.HandleFunc("/shortener", h.GetAllShorteners).Methods("GET")
	router.HandleFunc("/shortener/{slug}", h.GetShortener).Methods("GET")

	log.Println("API is running on port " + PORT)
	http.ListenAndServe(PORT, router)
}
