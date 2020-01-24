package main

import "fmt"

func main() {
	var chicken = make(map[string]int)

	chicken["january"] = 1

	fmt.Println(chicken)
}
