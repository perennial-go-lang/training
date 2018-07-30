package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/babulal107/training/Assignment/09_website/database"
	"github.com/gorilla/mux"
)

const port = 8787

var Route = mux.NewRouter()

func main() {
	Compile()
	Routes()
	database.Connect()
	http.Handle("/", Route)
	log.Printf("Starting server at port : %v", port)
	defer database.Db.Close()
	log.Fatalln(fmt.Println(http.ListenAndServe(fmt.Sprintf(":%v", port), Route)))
}
