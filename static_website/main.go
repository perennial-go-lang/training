package main

import (
	"log"
	"net/http"
)

func main() {
	Compile()
	filesHandler := http.FileServer(http.Dir("./public"))
	http.Handle("/", http.StripPrefix("/", filesHandler))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
