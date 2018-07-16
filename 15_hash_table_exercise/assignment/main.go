package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type end = []string
type aaa map[int]end

var bucket = [8]aaa{}

var occurance int32

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/74/74-0.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	var words int
	var word string

	// type aaa map[int]end
	// bucket := [8]aaa{}
	bucketCount := make([]int, 8)

	var el aaa
	for scanner.Scan() {
		words++
		word = scanner.Text()
		bId := getBucketId(word, 8)
		bucketCount[bId]++
		if bucket[bId] == nil {
			el = aaa{}
		} else {
			el = bucket[bId]
		}
		el[int(word[0])] = append(el[int(word[0])], word)
		bucket[bId] = el
	}

	// fmt.Println(bucket[7], bucketCount[7], len(bucket[7]))
	var count [8]int
	for index, bid := range bucket {
		// l := 0
		for key, val := range bid {
			fmt.Println(string(key), len(val))
			count[index] = count[index] + len(val)
		}
	}
	fmt.Println(count)

	wordTosearch := ""
	fmt.Print("Enter word to Search: ")
	fmt.Scanln(&wordTosearch)
	fmt.Println("Word is", wordTosearch)
	fmt.Println(search(wordTosearch))
}

func getBucketId(word string, numberOfBuckets int) int {
	letter := int(word[0])
	return letter % numberOfBuckets
}

func search(word string) int32 {
	bId := getBucketId(word, 8)
	occurance = 0
	arraySize := int(len(bucket[bId][int(word[0])]) / 100)
	sema := make(chan bool)
	for k := 0; k <= arraySize; k++ {
		go func(k int) {
			for _, j := range bucket[bId][int(word[0])][k*100 : (k*100 + 100)] {
				if j == word {
					atomic.AddInt32(&occurance, 1)
				}
			}
			sema <- true
		}(k)
	}

	for m := 0; m <= arraySize; m++ {
		<-sema
	}

	return occurance
}
