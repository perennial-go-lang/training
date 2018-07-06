package main

import (
	"os"
)

func main() {
	_, err := os.Open("asds.txt")
	if err != nil {
		//fmt.Println("Error!", err)
		//log.Println("Error!", err)
		//log.Fatalln("Error!", err)
		panic(err)
	}
}
