package main

import (
	"fmt"
)

func main() {
	// Basic Switch Case
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Unknown")
	// }

	// Multiple Condition Switch Case
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Weekday")
	// }

	// Type Switch Case
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I'm an integer")
		case string:
			fmt.Println("I'm a string")
		case bool:
			fmt.Println("I'm a boolean")
		default:
			fmt.Printf("Unknown type")
		}
	}

	whoAmI(1.00)
}