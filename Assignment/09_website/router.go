package main

import (
	"net/http"

	"github.com/babulal107/training/Assignment/09_website/controllers/user"
)

func Routes() {

	Route.HandleFunc("/users", user.Index).Methods("GET")
	Route.HandleFunc("/users/{id}", user.Show).Methods("GET")
	Route.HandleFunc("/users", user.Create).Methods("POST")
	Route.HandleFunc("/users/{id}", user.Update).Methods("PUT")
	Route.HandleFunc("/users/{id}", user.Delete).Methods("DELETE")

	fs := http.FileServer(http.Dir("./public/"))
	Route.PathPrefix("/").Handler(http.StripPrefix("/", fs))
}
