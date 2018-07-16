package main

import (
	template2 "text/template"
	"log"
	"os"
)

var template *template2.Template

func init() {
	template = template2.Must(template2.ParseGlob("templates/*"))
}

func main() {
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	f1, err := os.Create("index2.html")
	if err != nil {
		log.Fatal(err)
	}

	err = template.ExecuteTemplate(f, "index.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = template.ExecuteTemplate(f, "here.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = template.ExecuteTemplate(f1, "there.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
