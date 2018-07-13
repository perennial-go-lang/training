package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
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

	bucket := make([]int, 8)
	for scanner.Scan() {
		words++
		word = scanner.Text()
		bId := getBucketId(word, 8)
		bucket[bId]++
	}
	fmt.Println(words, bucket)
}

func getBucketId(word string, numberOfBuckets int) int {
	letter := int(word[0])
	return letter % numberOfBuckets
}
