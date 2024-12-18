package main

import "fmt"

// for -> only construct for loops/iteration in Go

func main() {
	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for loop
	// for i := 1; i <= 3; i++ {
	// 	// break
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// range
	for i := range 3{
		fmt.Println(i+1)
	}
}