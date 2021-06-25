package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const one_predicate = "a"
const two_predicate = "b"

func succes_place(result string, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Value": result, "Status": "true", "Err": ""})
}

func get_variable_string_from_uri(variable_name string, w http.ResponseWriter, r *http.Request) []string {
	variable, ok := r.URL.Query()[variable_name]
	if !ok || len(variable[0]) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "missed"}`))
	}
	return variable
}

func get_integers_for_equalize(one_name string, second_name string, w http.ResponseWriter, r *http.Request) (int, int) {

	one := get_variable_string_from_uri(one_name, w, r)
	two := get_variable_string_from_uri(second_name, w, r)
	first_int, err := strconv.Atoi(one[0])
	second_int, err := strconv.Atoi(two[0])
	if err != nil {
		fmt.Println((err))
	}
	return first_int, second_int
}

func div(w http.ResponseWriter, r *http.Request) {
	first, second := get_integers_for_equalize(one_predicate, two_predicate, w, r)
	result := first / second
	succes_place(fmt.Sprint(result), w, r)
}

func add(w http.ResponseWriter, r *http.Request) {
	first, second := get_integers_for_equalize(one_predicate, two_predicate, w, r)
	result := first + second
	succes_place(fmt.Sprint(result), w, r)
}

func sub(w http.ResponseWriter, r *http.Request) {
	first, second := get_integers_for_equalize(one_predicate, two_predicate, w, r)
	result := first - second
	succes_place(fmt.Sprint(result), w, r)
}

func mul(w http.ResponseWriter, r *http.Request) {
	first, second := get_integers_for_equalize(one_predicate, two_predicate, w, r)
	result := first * second
	succes_place(fmt.Sprint(result), w, r)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/add/", add)
	router.HandleFunc("/api/div/", div)
	router.HandleFunc("/api/sub/", sub)
	router.HandleFunc("/api/mul/", mul)
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Println("listening on 8080")
	http.ListenAndServe(":8080", loggedRouter)
}
