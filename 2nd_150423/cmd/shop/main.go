package main

import "fmt"

func main() {
	var test int // 0 // create in stack
	test += 1
	fmt.Printf("T before: %v\n", test)
	fmt.Printf("T after: %v\n") // create in heap
}
