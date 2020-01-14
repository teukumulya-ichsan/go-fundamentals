package main

import "fmt"

func main() {
	// deklarasi secara horizontal
	var number = [4]int64{1, 2, 3, 4}

	fmt.Println("Jumlah Element \t\t", len(number))
	fmt.Println("Isi element \t", number)

	var nama = [4]string{
		"Mulia Ichsan",
		"Error",
	}

	fmt.Println("Jumlah Element \t\t", len(nama))
	fmt.Println("Isi element \t", nama)

}
