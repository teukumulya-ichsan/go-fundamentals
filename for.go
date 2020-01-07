package main

import "fmt"

func main() {
	// i := 1

	// for i <= 10 {
	// 	fmt.Println(i, "hello for")
	// 	i++
	// }

	// transfor into

	for i := 1; i <= 10; i++ {
		fmt.Println(i, "Hello World")
	}
	fmt.Println("selesai incerment")

	for i := 10; i >= 1; i-- {
		fmt.Println(i, "Hello World")
	}

	fmt.Println("Selesai decremetn")
}
