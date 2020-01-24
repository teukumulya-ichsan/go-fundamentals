package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = []string{"John", "Wick"}

	printMessage("Hello", name)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)

}
