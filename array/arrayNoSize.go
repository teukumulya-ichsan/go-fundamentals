package main

import "fmt"

func main() {
	var nama = [...]string{
		"Mulia Ichsan",
		"Error",
	}

	fmt.Println("Jumlah Element \t\t", len(nama))
	fmt.Println("Isi element \t", nama)
}
