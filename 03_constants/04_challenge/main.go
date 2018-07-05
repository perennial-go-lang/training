package main

import "fmt"

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fmt.Printf("1 KB = %v bytes\n", KB)
	fmt.Printf("1 MB = %v KB = %v bytes\n", MB/KB, MB)
	fmt.Printf("1 GB = %v MB = %v KB = %v bytes\n", GB/MB, GB/KB, GB)
	fmt.Printf("1 TB = %v GB = %v MB = %v KB = %v bytes\n", TB/GB, TB/MB, TB/KB, TB)
}
