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

func main() {
	d1 := []int{10,11,12}

	err := tpl.ExecuteTemplate(os.Stdout, "layout.gohtml", d1)
	if err != nil {
		log.Fatal(err)
	}
}