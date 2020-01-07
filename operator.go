package main

import "fmt"

func main() {
	fmt.Println(1 + 1)    // add
	fmt.Println(10 - 1)   // substraction
	fmt.Println(10 * 2)   // multiplication
	fmt.Println(10 / 3.0) // division
	fmt.Println(10 % 3)   // remainder

	fmt.Println("\n")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
