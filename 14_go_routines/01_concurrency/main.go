package main

import (
	"fmt"
	"program_metrics"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	metrics.Start()
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	metrics.End()
}

func foo() {
	for i := 0; i < 50; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 50; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}
