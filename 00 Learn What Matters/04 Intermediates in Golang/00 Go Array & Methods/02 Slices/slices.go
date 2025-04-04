package main

import "fmt"

// slice -> dynamic array
//  most used construct in Go
// usefull methods: append, copy, len, cap, delete, pop, push, shift, unshift
func main() {
	// uninitialized slice is nil
	// var nums[]int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 2, 10)
	// fmt.Println(nums != nil)
	// Capacity is max size of the slice
	fmt.Println(cap(nums))
}