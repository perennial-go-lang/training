package main

import "fmt"

func main() {
	const (
		_  = iota
		KB = 1 << (iota * 10)
		MB
		GB
		TB
	)

	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d | KB = %d bytes\n", MB/KB, MB)
	fmt.Printf("1 GB = %d | MB = %d | KB = %d bytes\n", GB/MB, GB/KB, GB)
	fmt.Printf("1 TB = %d | GB = %d | MB = %d | KB = %d bytes\n", TB/GB, TB/MB, TB/KB, TB)
}
