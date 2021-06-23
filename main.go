package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func succes_place(result string, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Value": result, "Status": "true", "Err": ""})
}

func get_variable_from_uri(variable_name string, w http.ResponseWriter, r *http.Request) []string {
	variable, ok := r.URL.Query()[variable_name]
	fmt.Println(reflect.TypeOf(variable))
	if !ok || len(variable[0]) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "missed"}`))
	}
	return variable
}

func return_variables_for_equalize(one_name string, second_name string, w http.ResponseWriter, r *http.Request) ([]string, []string) {
	one := get_variable_from_uri(one_name, w, r)
	two := get_variable_from_uri(second_name, w, r)
	return one, two
}

func div(w http.ResponseWriter, r *http.Request) {
	first, second := return_variables_for_equalize("a", "b", w, r)
	fmt.Println(second[0])
	//TODO add hanler for equalize two variales
	succes_place(fmt.Sprint(first[0]), w, r)
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
