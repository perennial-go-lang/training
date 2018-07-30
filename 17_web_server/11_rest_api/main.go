package main

import (
	"fmt"
	"log"
	"net/http"

	route "github.com/ankitkumbhar/training/17_web_server/11_rest_api/Route"
)

const port = 6767

func main() {
	r := route.Router()

	fmt.Printf("Starting server on port : %v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), r)
	if err != nil {
		log.Fatal(err)
	}
}
