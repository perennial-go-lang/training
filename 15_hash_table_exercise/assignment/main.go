package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

const NUMBER_OF_BUCKETS = 8

type aaa map[int][]string

var bucket = [NUMBER_OF_BUCKETS]aaa{}

var occurance int32
var bucketCount [NUMBER_OF_BUCKETS]int
var count [NUMBER_OF_BUCKETS]int

func main() {

	fmt.Println("**************************************************************")
	fmt.Println("*                Welcome to Word Sort                        *")
	fmt.Println("**************************************************************")

	fmt.Println("* Please enter file URL:                                     *")

	fileUrl := ""
	fmt.Scanln(&fileUrl)
	res, err := http.Get(fileUrl)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	scanner.Split(bufio.ScanWords)

	var words int
	var word string

	var el aaa

	for scanner.Scan() {
		word = scanner.Text()
		bId := getBucketId(word, NUMBER_OF_BUCKETS)

		words++
		bucketCount[bId]++
		if bucket[bId] == nil {
			el = aaa{}
		} else {
			el = bucket[bId]
		}
		el[int(word[0])] = append(el[int(word[0])], word)
		bucket[bId] = el
	}

	// time.Sleep(5 * time.Second)
	for index, bid := range bucket {
		for _, val := range bid {
			// fmt.Println(string(key), len(val))
			count[index] = count[index] + len(val)
		}
	}
	fmt.Printf("Total words found: %d\n", words)

	doMore := true
	for doMore == true {
		fmt.Println("**************************************************************")
		fmt.Println("* Please choose option:                                      *")
		fmt.Println("* 1. Search by word            2.Count by initial            *")
		fmt.Println("* 3. Bucket wise count                                       *")
		fmt.Println("**************************************************************")

		option := 0
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("* Enter word: 												  *")
			wordTosearch := ""
			fmt.Scanln(&wordTosearch)
			fmt.Println("Word is", wordTosearch)
			fmt.Println("We found ", search(wordTosearch), " occurarences of word: ", wordTosearch)
		case 2:
			fmt.Println("* Enter initial: 										      *")
			wordTosearch := ""
			fmt.Scanln(&wordTosearch)
			bId := getBucketId(wordTosearch, NUMBER_OF_BUCKETS)
			fmt.Println("Count is", len(bucket[bId][int(wordTosearch[0])]))
		case 3:
			showBucketWise()
		default:
			fmt.Println("Please enter write choice")
		}
		fmt.Println("Do you want to do more? Y:yes, N: No")
		wantMore := "Y"
		fmt.Scanln(&wantMore)
		correctChoice := false
		for correctChoice == false {
			switch wantMore {
			case "Y":
				doMore = true
				correctChoice = true
			case "N":
				doMore = false
				correctChoice = true
			default:
				fmt.Println("Please enter write choice Y/N")
				correctChoice = false
				doMore = false
				fmt.Println("Do you want to do more? Y:yes, N: No")
				wantMore := "Y"
				fmt.Scanln(&wantMore)
			}
		}
	}
}

func getBucketId(word string, numberOfBuckets int) int {
	letter := int(word[0])
	return letter % numberOfBuckets
}

func search(word string) int32 {
	bId := getBucketId(word, NUMBER_OF_BUCKETS)
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

func showBucketWise() {
	for i := 0; i < NUMBER_OF_BUCKETS; i++ {
		fmt.Printf("Bucket %d : %d\n", i, bucketCount[i])
	}
}
