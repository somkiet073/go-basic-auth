package main

import (
	"net/http"

	"go-basic-auth/api"
	"go-basic-auth/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/api/demo/demo1", middlewares.BasicAuthMiddleware(http.HandlerFunc(api.Demo1API))).Methods("GET")
	r.HandleFunc("/api/demo/demo2", api.Demo2API).Methods("GET")

	if err := http.ListenAndServe(":8181", r); err != nil {
		panic(err)
	}
}
