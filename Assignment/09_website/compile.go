package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func Compile() {
	Data := []int{10, 11, 12}

	file, err := os.Create("./public/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(file, "layout.gohtml", Data)
	if err != nil {
		log.Fatal(err)
	}
}
