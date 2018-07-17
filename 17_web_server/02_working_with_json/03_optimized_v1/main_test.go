package main

import (
	"testing"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func BenchmarkHandlerMarshal(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	message := Msg{"Hello, World!"}
	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(message)
		fmt.Fprint(writer, string(data))
	}
}

func BenchmarkHandlerEncoder(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	message := Msg{"Hello, World!"}
	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(message)
	}
}

func BenchmarkHandlerEncoderReference(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	message := Msg{"Hello, World!"}
	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(&message)
	}
}