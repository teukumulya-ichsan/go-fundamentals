package main

import "fmt"

func main() {
	var data = [2][2]string{
		[2]string{"Mulia", "ichsan"},
		[2]string{"Geubrina", "Melati"},
	}

	//atau
	var data2 = [2][2]string{
		{"Mulia", "ichsan"},
		{"Geubrina", "Melati"},
	}

	fmt.Println("jumlah element", len(data))
	fmt.Println(" element", data)

	fmt.Println("jumlah element", len(data2))
	fmt.Println(" element", data2)
}
