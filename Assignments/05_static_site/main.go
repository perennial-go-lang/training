package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	filesHandler := http.FileServer(http.Dir("./public"))
	http.Handle("/home/", http.StripPrefix("/home/", filesHandler))

	log.Printf("Starting server at port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
