package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println("Angka", i)
		i++

		if i == 5 {
			break
		}
	}

}
