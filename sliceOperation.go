package main

import "fmt"

func main() {
	// slice := []string{
	// 	"geby", "icut", "hom",
	// }

	slice := make([]string, 3)
	slice[0] = "geby"
	slice[1] = "melati"
	slice[2] = "hom"

	slice2 := append(slice, "budy", "nugraha")

	fmt.Println(slice)
	fmt.Println(slice2)

	slice[2] = "Mae"

	fmt.Println(slice)
	fmt.Println(slice2)

	//copy slice 1 to slice 3
	slice3 := make([]string, 2)
	copy(slice3, slice) // copy hanyaa akan mengcopy seberapa muat element saja
	fmt.Println(slice3)
}
