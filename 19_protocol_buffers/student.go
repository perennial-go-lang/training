package main

import (
	"os"
	"encoding/json"
	"fmt"
	"github.com/perennial-go-training/training/19_protocol_buffers/json"
)

func main () {
	s1 := student.Student{
		"a",
		"b",
		3,
	}

	var s2 student.Student

	f, _ := os.Create("./student.json")
	encoder := json.NewEncoder(f)
	encoder.Encode(&s1)

	f, _ = os.Open("student.json")
	decoder := json.NewDecoder(f)
	decoder.Decode(&s2)

	fmt.Println(s2)
}
