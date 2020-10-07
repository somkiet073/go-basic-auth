package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Method use OK")
}

func postData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Method postData use OK ")
}

func putData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	id := mux.Vars(r)
	fmt.Fprintf(w, "putDataid: %v", id["id"])
}

func delData(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	id := mux.Vars(r)
	fmt.Fprintf(w, "DelDataid: %v", id["id"])
}

func handleAPI() {
	r := mux.NewRouter()
	r.HandleFunc("/user", homeHandler).Methods("GET")
	r.HandleFunc("/user", postData).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}", putData).Methods("PUT")
	r.HandleFunc("/user/{id:[0-9]+}", delData).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}

func main() {
	handleAPI()
}
