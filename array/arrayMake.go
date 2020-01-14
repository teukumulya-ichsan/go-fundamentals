package main

import "fmt"

func main() {

	var nama = make([]string, 2)

	nama[0] = "Geubrina"
	nama[1] = "Melati"

	for i, val := range nama {
		fmt.Println(i, val)
	}
	fmt.Println(nama)

}
