package main

import "fmt"

func main() {

	var names [3]string
	fmt.Println(names)

	var girl [3]string
	girl[0] = "geby"
	girl[1] = "icut"

	// fmt.Println(girl[0])
	// fmt.Println(girl[1])

	for i := 0; i < len(girl); i++ {
		fmt.Println(girl[i])
	}

	// foreach array on go

	for _, value := range girl {
		// fmt.Println("index", index, "=", value)
		fmt.Println(value)
	}

	// inline array declaration
	lang := [3]string{
		"Go", "Javascript", "Ruby on Rails",
	}

	for _, value := range lang {
		fmt.Println(value)
	}

	fmt.Println(lang)
}
