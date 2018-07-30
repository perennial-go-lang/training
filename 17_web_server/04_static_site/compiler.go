package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func Compile() {
	d1 := []int{10,11,12,14}

	file, err := os.Create("./public/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(file, "layout.gohtml", d1)
	if err != nil {
		log.Fatal(err)
	}
}