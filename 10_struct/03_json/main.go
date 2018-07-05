package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FName string `json:"FirstName"`
	LName string `json:"LastName"`
	age int
}

func main() {
	p1 := person{"A", "B", 22}
	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))
	fmt.Printf("%T\n", p1Json)

	var p2 person
	p2Json := []byte(`{"FirstName":"X", "LastName": "Y", "Age": 33 }`)
	_ = json.Unmarshal(p2Json, &p2)

	fmt.Println(p2)
	fmt.Printf("%T\n", p2)
}