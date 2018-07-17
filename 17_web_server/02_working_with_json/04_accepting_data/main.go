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
	var incomingMessage Msg

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incomingMessage)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	message := Msg{Data: "Hello, " + incomingMessage.Data + "!"}

	encoder := json.NewEncoder(w)
	encoder.Encode(&message)
}
