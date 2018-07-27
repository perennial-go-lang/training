package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type MyHandler struct{}

func main() {
	http.HandleFunc("/foo", fooHandler)
	http.Handle("/notfound/", http.NotFoundHandler())

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

	s := &http.Server{
		Addr:           ":8082",
		Handler:        nil,
		ReadTimeout:    1 * time.Second,
		WriteTimeout:   1 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
