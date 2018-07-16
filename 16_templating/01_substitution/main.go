package main

import (
	"os"
	"log"
	"io"
	"strings"
)

func  main() {
	name := os.Args[1]

	template := `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Welcome to Go Templating</title>
		</head>
		<body>
			<h1>` + name + `</h1>
		</body>
	</html>
	`
	f, err := os.Create("indexnew.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	io.Copy(f, strings.NewReader(template))
}
