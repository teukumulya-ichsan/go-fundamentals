package main

import "fmt"

func main() {
	var data = [...]string{
		"Mulia", "ichsan", "sss",
	}

	for i, datas := range data {
		fmt.Printf("elemen %d : %s\n", i, datas)
	}
}
