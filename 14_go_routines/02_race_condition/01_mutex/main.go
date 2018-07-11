package main

import (
	"fmt"
	"sync"
	"time"
	"program_metrics"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func main() {
	metrics.Start()
	wg.Add(3)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	go incrementor("New: ")
	wg.Wait()
	metrics.End()
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		mutex.Lock()
		time.Sleep(5 * time.Millisecond)
		fmt.Println(s, counter)
		counter++
		mutex.Unlock()
	}
	wg.Done()
}
