package main

import (
	"fmt"
)

func main() {

	fmt.Println("Sum : ", sum(2, 3, 4))

}

func sum(number ...int) (result int) {
	// var total int = 0
	for _, value := range number {
		result += value
	}

	return
}
