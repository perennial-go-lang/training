package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohitpavaskar/training/api/models"
)

const port = 8080

var Route *mux.Router

func main() {
	Route = mux.NewRouter()
	BootstrapRoutes()
	models.Connect()
	defer models.Close()
	http.Handle("/", Route)
	log.Println("Starting server ad port: ", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", port), Route))
}
