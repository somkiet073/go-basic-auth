package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// เขียนแบบ แยก Method
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

// End เขียนแบบ แยก Method

func main() {
	r := mux.NewRouter()

	// เขียนแบบแยก
	r.HandleFunc("/user", homeHandler).Methods("GET")
	r.HandleFunc("/user", postData).Methods("POST")
	r.HandleFunc("/user/{id:[0-9]+}", putData).Methods("PUT")
	r.HandleFunc("/user/{id:[0-9]+}", delData).Methods("DELETE")

	// เขียนแบบรวม
	r.HandleFunc("/user2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Http Get Ok")
	}).Methods("GET")

	r.HandleFunc("/user2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Http Post OK")
	}).Methods("POST")

	r.HandleFunc("/user2/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)
		fmt.Fprintf(w, "Http Put OK id: %v", id["id"])
	}).Methods("PUT")

	r.HandleFunc("/user2/{name:[a-z]+}", func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)
		fmt.Fprintf(w, "Http Delete OK id: %v", name["name"])
	}).Methods("DELETE")

	// run serve :8081
	http.ListenAndServe(":8081", r)

}
