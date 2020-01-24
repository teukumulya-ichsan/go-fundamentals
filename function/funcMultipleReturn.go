package main

import "fmt"

func main() {

	var luas, keliling = luasKeliling(2, 3, 4)
	fmt.Println("Luas : ", luas)
	fmt.Println("Keliling : ", keliling)
}

func luasKeliling(x, y, z int) (luas, keliling int) {

	luas = x * y * z
	keliling = x + y + z

	return luas, keliling
}
