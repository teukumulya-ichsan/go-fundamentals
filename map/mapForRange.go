package main

import (
	"fmt"
)

func main() {
	var lang = map[string]string{"js": "Javascript", "Go": "Golang"}

	for key, value := range lang {
		fmt.Println(key, "\t:", value)
	}
}
