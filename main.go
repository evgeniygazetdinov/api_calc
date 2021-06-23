package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func div(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"apple": 1})

}
func add(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(1)
}
func sub(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(1)
}
func mul(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(1)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/add/", add)
	router.HandleFunc("/api/div/", div)
	router.HandleFunc("/sub/", sub)
	router.HandleFunc("/mul/", mul)
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Println("listening on 8080")
	http.ListenAndServe(":8080", loggedRouter)
}
