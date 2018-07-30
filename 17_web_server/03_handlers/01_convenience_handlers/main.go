package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port = 6767

type MyHandler struct{}

func main() {
	filesHandler := http.FileServer(http.Dir("/home/ajitem"))
	http.Handle("/files/", http.StripPrefix("/files/", filesHandler))

	http.Handle("/notfound/", http.NotFoundHandler())

	http.Handle("/folders/", http.RedirectHandler("/files/", http.StatusMovedPermanently))

	timedHandler := &MyHandler{}
	http.Handle("/time/", http.TimeoutHandler(timedHandler, time.Second, "Process Timed Out"))

	log.Print("Starting server at port %v", 6767)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Hello, World!")
}
