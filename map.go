package main

import "fmt"

func main() {

	girl := make(map[string]string) //make a map

	girl["01"] = "geby" // append geby to keys 01
	girl["02"] = "melati"
	fmt.Println(girl)

	fmt.Println(girl["01"]) // print girls with keys 01
	fmt.Println(girl["012"])

	for keys, name := range girl {
		fmt.Println("Keys", keys, "name is", name)
	}

	// make direct map
	lang := map[string]string{
		"Go":  "Golang",
		"JS":  "Javascript",
		"RoR": "Ruby on Rails",
	}

	fmt.Println(lang)

	fmt.Println(lang["Go"])
}
