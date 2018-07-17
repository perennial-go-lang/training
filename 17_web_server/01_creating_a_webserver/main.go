package main

import (
	"net/http"
	"fmt"
	"log"
)

const port = 6767

func main() {
	http.HandleFunc("/helloworld", myHandler)
	log.Printf("Starting server at port %v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
