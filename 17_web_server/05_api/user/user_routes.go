package user

import (
	"github.com/perennial-go-training/training/17_web_server/05_api/router"
	"log"
)

func bootstrapRoutes() {
	if router.Router == nil {
		log.Fatal("Router not initialized")
	}

	s := router.Router.PathPrefix("/users").Subrouter()

	s.HandleFunc("", CreateUser).Methods("POST")
	s.HandleFunc("", GetUsers).Methods("GET")
	s.HandleFunc("/{id}", GetUser).Methods("GET")
	s.HandleFunc("/{id}", UpdateUser).Methods("PUT")
	s.HandleFunc("/{id}", DeleteUser).Methods("DELETE")
}
