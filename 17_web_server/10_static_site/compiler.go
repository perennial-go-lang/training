package main

import (
	"log"
	"os"
	"text/template"
)

type student struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	s1 := []student{
		{
			"FirstStudent",
			20,
		},
		{
			"SecondStudent",
			18,
		},
	}

	file, err := os.Create("public/index.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(file, "first.gohtml", s1)
	if err != nil {
		log.Fatal(err)
	}
}
