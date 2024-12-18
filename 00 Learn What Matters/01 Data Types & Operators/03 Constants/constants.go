package main

import "fmt"

const pi = 3.14

func main() {
	// := cannot be used outside functions
	// const pi = 3.14 cannot be redeclared

	fmt.Println(pi)

	// constant grouping
	const (
		port = 200
		host = "localhost"
	)

	// port = 300 // cannot be redeclared

	fmt.Println(port, host)
}