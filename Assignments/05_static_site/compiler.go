package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	d1 := []int{10, 11, 12, 14, 16}

	file, err := os.Create("./public/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(file, "layout.gohtml", d1)
	if err != nil {
		log.Fatal(err)
	}

	// err := tpl.ExecuteTemplate(os.Stdout, "layout.gohtml", d1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
