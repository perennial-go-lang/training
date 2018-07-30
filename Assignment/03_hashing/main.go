package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/1342/1342-0.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	var words int
	var word string
	type end = []string
	type aaa map[int]end
	bucket := [8]aaa{}
	bucketCount := make([]int, 8)

	var el aaa

	for scanner.Scan() {
		words++
		word = scanner.Text()
		bID := getBucketID(word, 8)
		bucketCount[bID]++
		if el == nil {
			el = aaa{}
		}
		el[int(word[0])] = append(el[int(word[0])], word)
		bucket[bID] = el
	}
	// fmt.Println(bucket[7], bucketCount[7], len(bucket[7]))
	for _, bid := range bucket {
		for key, val := range bid {
			fmt.Println(string(key), len(val))
		}
	}
}

func getBucketID(word string, numberOfBuckets int) int {
	letter := int(word[0])
	return letter % numberOfBuckets
}
