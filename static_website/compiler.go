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
	file, err := os.Create("./public/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(file, "layout.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

}
