package main

import (
	"github.com/perennial-go-training/training/19_protocol_buffers/proto/studentpb"
	"github.com/golang/protobuf/proto"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	s1 := studentpb.Student{
		FirstName: "a",
		LastName: "b",
		Age: 3,
	}

	var s2 studentpb.Student

	bytes, err := proto.Marshal(&s1)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("student.bin", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err = ioutil.ReadFile("student.bin")
	proto.Unmarshal(bytes, &s2)

	fmt.Println(s2.GetFirstName())
	fmt.Println(s2.GetLastName())
	fmt.Println(s2.GetAge())
}

