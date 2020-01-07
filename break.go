package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			continue // if i is ganjil, is skipped
		}

		fmt.Println(i, "Genap")

		if i == 50 {
			break // if i has reached 50, looping is stopped
		}

	}
}
