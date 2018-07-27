package main

import (
	"fmt"
	"net/http"
)

const port = 6767

func main() {
	filesHandler := http.FileServer(http.Dir("./public"))
	http.Handle("/files/", http.StripPrefix("/files/", filesHandler))

	fmt.Printf("Starting server on port : %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
