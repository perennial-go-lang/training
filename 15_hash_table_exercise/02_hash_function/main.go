package main

import "fmt"

func main() {
	word := "GoLang"
	bucketId := getBucketId(word, 8)

	fmt.Println("BucketID:", bucketId)
}

func getBucketId(word string, numberOfBuckets int) int {
	letter := int(word[0])
	return letter % numberOfBuckets
}
