package main

import (
	"fmt"
	"log"
	"net/http"

	route "github.com/soniraju/training/Assignments/06_basic_api/Router"
)

const port = 5656

func main() {
	r := route.Router()

	fmt.Printf("Starting server on port : %v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), r)
	if err != nil {
		log.Fatal(err)
	}
}
