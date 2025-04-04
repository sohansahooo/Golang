package main

import "fmt"

func main() {
	// var num[4]int
	// fmt.Println(len(num))
	// num[0] = 1
	// fmt.Println(num)

	// var vals[3]bool
	// vals[0] = true
	// fmt.Println(vals)

	// var name[3]string
	// name[0] = "John"
	// // name[1] = "Doe"
	// // name[2] = "Smith"
	// fmt.Println(len(name),name)

	// In zeroed state,
	// int: 0, string: "", bool: false

	// To declare and initialize an array
	// nums := [3]int{1, 2, 3}
	// fmt.Println(nums)

	//  2D Array 
	nums := [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums)
	// -- if fixed size bcuz, it's predictable
	// -- Memory optimization
	// -- Constant time access
	
}