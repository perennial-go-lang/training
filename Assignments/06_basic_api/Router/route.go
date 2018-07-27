package route

import (
	"github.com/gorilla/mux"
	"github.com/soniraju/training/Assignments/06_basic_api/Employee"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/employees", Employee.Index).Methods("GET")
	r.HandleFunc("/employees", Employee.Store).Methods("POST")
	r.HandleFunc("/employees/{id:[0-9]+}", Employee.Show).Methods("GET")
	r.HandleFunc("/employees/{id:[0-9]+}", Employee.Destroy).Methods("DELETE")
	r.HandleFunc("/employees/{id:[0-9]+}", Employee.Update).Methods("PUT")
	return r
}
