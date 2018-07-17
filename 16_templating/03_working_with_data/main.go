package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var tpl *template.Template

type Student struct {
	Name       string
	Age        int
	Percentage float64
}

var fm = template.FuncMap{
	"lc":   strings.ToLower,
	"per":  formatPercentage,
	"itos": intToString,
	"age":  formatAge,
}

func formatPercentage(percentage float64) string {
	return strconv.FormatFloat(percentage, 'f', 2, 64) + "%"
}

func intToString(number int) string {
	return strconv.FormatInt(int64(number), 10)
}
func formatAge(age string) string {
	return age + " Years Old"
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func main() {
	//err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "World!")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//s := []string{"Ankit", "Babulal", "Rohit", "Soni"}
	//err = tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", s)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//m := map[int]string{
	//	0: "Ankit",
	//	1: "Babulal",
	//	2: "Rohit",
	//	3: "Soni",
	//}
	//err = tpl.ExecuteTemplate(os.Stdout, "map.gohtml", m)
	//if err != nil {
	//	log.Fatal(err)
	//}

	students := []Student{
		{
			"B",
			20,
			50,
		},
		{
			"A",
			19,
			60,
		},
		{
			"C",
			18,
			55,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", students)
	if err != nil {
		log.Fatal(err)
	}
}
