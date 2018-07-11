package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"program_metrics"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	metrics.Start()
	wg.Add(3)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	go incrementor("New: ")
	wg.Wait()
	fmt.Println(counter)
	metrics.End()
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		fmt.Println(s, atomic.LoadInt64(&counter))
		atomic.AddInt64(&counter, 1)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}
