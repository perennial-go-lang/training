package main

import (
	"fmt"
	"sync"
)

func main() {
	//c := onDemandGenerator()
	//for i := 0; i < 100; i++ {
	//	fmt.Println(<-c)
	//}

	// PIPELINE
	//for n := range cube(square(generator(1, 5))) {
	//	fmt.Println(n)
	//}

	input := generator(1, 20)

	// FAN OUT
	sq := square(input)
	cu := cube(input)

	// FAN IN
	for s:= range merge(sq, cu) {
		fmt.Println(s)
	}
}

//func onDemandGenerator() <-chan int {
//	out := make(chan int)
//	var i int
//	go func() {
//		for {
//			i++
//			out <- i
//		}
//		close(out)
//	}()
//	return out
//}

func generator(start, end int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; i <= end ; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func cube(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n * n
		}
		close(out)
	}()
	return out
}

func merge(cs... <-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}