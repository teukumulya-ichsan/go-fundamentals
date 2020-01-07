package main

import "fmt"

func main() {
	// inline array declaration
	lang := [4]string{
		"Go", "Javascript", "Ruby on Rails", "TypeScript",
	}

	fmt.Println(lang)

	slice1 := lang[1:3] // return data from index 1 to data 3 on array of lang
	fmt.Println(slice1)

	slice2 := lang[:2] // return data until 2nd data
	fmt.Println(slice2)

	slice3 := lang[1:] // return data from index 0 until last
	fmt.Println(slice3)

}
