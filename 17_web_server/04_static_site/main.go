package main

import (
	"log"
	"github.com/fsnotify/fsnotify"
	"net/http"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func () {
		filesHandler := http.FileServer(http.Dir("./public"))
		http.Handle("/", http.StripPrefix("/", filesHandler))
		log.Fatal(http.ListenAndServe(":8080",nil))
	}()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("EVENT!!", event)
				Compile()
			case err := <-watcher.Errors:
				log.Fatal("ERROR", err)
			}
		}
	}()

	err = watcher.Add("./templates")
	if err != nil {
		log.Fatal(err)
	}

	<-done
}