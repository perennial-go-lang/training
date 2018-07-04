package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d | KB = %d bytes \n", MB/KB, MB)
	fmt.Printf("1 GB = %d | MB = %d | KB = %d bytes \n", GB/MB, GB/KB, GB)
	fmt.Printf("1 TB = %d | TB = %d | MB = %d | KB = %d bytes \n", TB/GB, TB/MB, TB/KB, TB)
}
