package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

const port = 6767

type Msg struct {
	Data string `json:"data"`
}

func main() {
	http.HandleFunc("/helloworld", myHandler)
	log.Printf("Starting server on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	message := Msg{Data: "Hello, World!"}

	encoder := json.NewEncoder(w)
	encoder.Encode(message)
}
