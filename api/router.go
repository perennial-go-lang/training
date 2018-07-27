package main

import "github.com/rohitpavaskar/training/api/controllers/post"

func BootstrapRoutes() {
	// Route.HandleFunc("/posts", post.Index).Methods("GET")
	posts := Route.PathPrefix("/posts").Subrouter()
	posts.HandleFunc("", post.Index).Methods("GET")
	posts.HandleFunc("/{id}", post.Show).Methods("GET")
	posts.HandleFunc("/{id}", post.Update).Methods("PUT")
	posts.HandleFunc("", post.Store).Methods("POST")
	posts.HandleFunc("/{id}", post.Destroy).Methods("DELETE")
}
