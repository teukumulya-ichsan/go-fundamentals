package main

import "fmt"

func main() {

	lang := map[string]string{
		"Go":  "Golang",
		"JS":  "Javascript",
		"RoR": "Ruby on Rails",
	}
	fmt.Println(lang) // print all lang map

	// declare map with value as map
	stack := map[string]map[string]string{
		"Go": map[string]string{
			"name":   "Golang",
			"salary": "> Rp. 15.000.000",
		},
		"JS": map[string]string{
			"name":   "Javascript",
			"salary": "> Rp. 12.000.000",
		},
		"RoR": map[string]string{
			"name":   "Ruby on Rails",
			"salary": "> Rp. 10.000.000",
		},
	}

	fmt.Println(stack)                 // print all on stack map
	fmt.Println(stack["JS"])           // print all on stack name with keys as "JS"
	fmt.Println(stack["JS"]["salary"]) // print salary on stack name with keys as "JS"

	// delete element
	delete(stack, "RoR")

	// adding element
	stack["PHP"] =
		map[string]string{
			"name":   "PHP",
			"salary": "> Rp. 5.000.000",
		}

	fmt.Println(stack) // print all on stack map after delete and add
}
