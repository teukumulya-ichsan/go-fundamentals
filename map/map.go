package main

import (
	"fmt"
)

func main() {
	var chicken = map[string]int{}

	chicken["january"] = 20

	fmt.Println("Secara eksplisit: \n", "January", chicken["january"])

	// direct
	var lang = map[string]string{"js": "Javascript", "Go": "Golang"}
	fmt.Println("Isi", lang)
	fmt.Println(lang["js"])

}
