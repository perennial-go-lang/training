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

	bucket := make([]int, 12)
	for scanner.Scan() {
		words++
		word = scanner.Text()
		bId := getBucketId(word, 12)
		bucket[bId]++
	}
	fmt.Println(words, bucket)
}

func getBucketId(word string, numberOfBuckets int) int {
	var sum int
	for _, char := range word  {
		sum += int(char)
	}
	return sum % numberOfBuckets
}
