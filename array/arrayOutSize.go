package main

import "fmt"

func main() {
	var name [4]string
	name[0] = "trafalgar"
	name[1] = "d"
	name[2] = "walter"
	name[3] = "law"
	name[4] = "ok" // out of bound size of array

	// for _, val := range name {
	// 	fmt.Print(val + " ")
	// }
	// fmt.Println()

	fmt.Println(name[0], name[1], name[2], name[3])
}
