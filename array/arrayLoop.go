package main

import "fmt"

func main() {
	var data = [2][]string{
		{"Mulia", "ichsan", "sss"},
		{"Geubrina", "Melati"},
	}

	for i := 0; i < len(data); i++ {
		for j := i; j < len(data[i]); j++ {
			fmt.Print(data[i][j] + " ")
		}
		fmt.Println()
	}
}
