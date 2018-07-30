package main

import (
	"testing"
	"github.com/perennial-go-training/training/19_protocol_buffers/json"
	"github.com/perennial-go-training/training/19_protocol_buffers/proto/studentpb"
	"encoding/json"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"os"
)

var s1, s2 student.Student
var s3, s4 studentpb.Student

func init() {
	s1 = student.Student{
		"a",
		"b",
		3,
	}

	s3 = studentpb.Student{
		FirstName: "a",
		LastName: "b",
		Age: 3,
	}
}

func BenchmarkToJson(b *testing.B) {
	b.ResetTimer()

	encoder := json.NewEncoder(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		encoder.Encode(&s1)
	}
}

func BenchmarkToProto(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		proto.Marshal(&s3)
	}
}

func BenchmarkFromJson(b *testing.B) {
	f, _ := os.Open("student.json")

	for i := 0; i < b.N; i++ {
		decoder := json.NewDecoder(f)
		decoder.Decode(&s2)
	}
}

func BenchmarkFromProto(b *testing.B) {
	bytes, _ := ioutil.ReadFile("student.bin")

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bytes, &s4)
	}
}
