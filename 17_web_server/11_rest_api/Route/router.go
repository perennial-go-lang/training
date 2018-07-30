package route

import (
	"github.com/ankitkumbhar/training/17_web_server/11_rest_api/Controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", Controller.Index).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", Controller.Show).Methods("GET")
	r.HandleFunc("/users", Controller.Store).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", Controller.Destroy).Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}", Controller.Update).Methods("PUT")
	return r
}
